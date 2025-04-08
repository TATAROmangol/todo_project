package auth

import (
	"context"
	"time"
	"todo/pkg/logger"

	authpb "todo/pkg/grpc/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	conn   *grpc.ClientConn
	client authpb.AuthClient
)

type AuthClient struct {
	cfg Config
}

func NewAuthClient(cfg Config) *AuthClient {
	return &AuthClient{cfg}
}

func (c *AuthClient) connect(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	con, err := grpc.NewClient(
		c.cfg.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logger.GetFromCtx(ctx).ErrorContext(ctx, ErrConnectNginx, err)
		return err
	}

	logger.GetFromCtx(ctx).InfoContext(ctx, "Listen auth", "path", c.cfg.Address)

	conn = con
	client = authpb.NewAuthClient(conn)
	return nil
}

func (c *AuthClient) GetId(ctx context.Context, token string) (int, error) {
	if conn == nil {
		if err := c.connect(ctx); err != nil {
			logger.GetFromCtx(ctx).ErrorContext(ctx, ErrConnectNginx, err)
			return -1, err
		}
	}

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	resp, err := client.GetId(ctx, &authpb.JWTRequest{Token: token})
	if err != nil{
		logger.GetFromCtx(ctx).ErrorContext(ctx, ErrGetIdGRPC, err)
		return 0, err
	}
	return int(resp.GetId()), nil
}

func (c *AuthClient) Close() error {
	return conn.Close()
}

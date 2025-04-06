package auth

import (
	"context"
	"time"
	"todo/pkg/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	authpb "todo/pkg/grpc/auth"
)

type AuthClient struct {
	conn   *grpc.ClientConn
	client authpb.AuthClient
}

func NewAuthClient(ctx context.Context, cfg Config) (*AuthClient, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	conn, err := grpc.NewClient(
		cfg.Address, 
		grpc.WithTransportCredentials(insecure.NewCredentials()), 
	)
	if err != nil {
		logger.GetFromCtx(ctx).ErrorContext(ctx, "Failed to connect to Nginx", "error", err)
		return nil, err
	}

	logger.GetFromCtx(ctx).InfoContext(ctx, "Listen auth", "path",cfg.Address)

	return &AuthClient{
		conn:   conn,
		client: authpb.NewAuthClient(conn),
	}, nil
}

func (c *AuthClient) GetId(ctx context.Context, token string) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	resp, err := c.client.GetId(ctx, &authpb.JWTRequest{Token: token})
	if err != nil {
		return 0, err
	}
	return int(resp.Id), nil
}

func (c *AuthClient) Close() error {
	return c.conn.Close()
}
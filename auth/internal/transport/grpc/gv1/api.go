package gv1

import (
	ssov1 "auth/pkg/grpc/auth"
	"auth/pkg/logger"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

type Auth interface {
	GetId(string) (int, error)
}

type Api struct {
	ssov1.UnimplementedAuthServer
	service Auth
}

func Register(gRPCServer *grpc.Server, service Auth) {
	ssov1.RegisterAuthServer(gRPCServer, &Api{service: service})
}

func (s *Api) GetId(
	ctx context.Context,
	in *ssov1.JWTRequest,
) (*ssov1.IdResponse, error) {
	if in.GetToken() == ""{
		return nil, fmt.Errorf("")
	}

	id, err := s.service.GetId(in.Token)
	if err != nil{
		return nil, err
	}

	logger.GetFromCtx(ctx).InfoContext(ctx, "get id", "id", id)

	return &ssov1.IdResponse{Id: int64(id)}, nil
}

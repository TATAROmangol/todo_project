package gv1

import (
	"auth/internal/errors"
	ssov1 "auth/pkg/grpc/auth"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

//go:generate mockery --all --output=./mocks

type Auth interface {
	Register(string, string) (string, error)
	Login(string, string) (string, error)
}

type Api struct {
	ssov1.UnimplementedAuthServer
	service Auth
}

func Register(gRPCServer *grpc.Server, service Auth) {
	ssov1.RegisterAuthServer(gRPCServer, &Api{service: service})
}

func (s *Api) Login(
	ctx context.Context,
	in *ssov1.LoginRequest,
) (*ssov1.TokenResponse, error) {
	if in.GetLogin() == ""{
		return nil, fmt.Errorf("")
	}
	if in.GetPassword() == ""{
		return nil, fmt.Errorf("")
	}

	token, err := s.service.Login(in.Login, in.Password)
	if err != nil{
		return nil, err
	}

	return &ssov1.TokenResponse{Token: token}, nil
}

func (s *Api) Register(
	ctx context.Context,
	in *ssov1.RegisterRequest,
) (*ssov1.TokenResponse, error) {
	if in.GetLogin() == ""{
		return nil, errors.ErrInvalidLogin
	}
	if in.GetPassword() == ""{
		return nil, errors.ErrInvalidPassword
	}

	token, err := s.service.Register(in.Login, in.Password)
	if err != nil{
		return nil, err
	}

	return &ssov1.TokenResponse{Token: token}, nil
}

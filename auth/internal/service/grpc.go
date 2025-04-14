package service

import (
	"auth/pkg/logger"
	"context"
)
//go:generate mockgen -destination=./mock/mock_grpc.go -package=mock -source=grpc.go

type JWTValidator interface{
	GetId(string) (int, error)
}

type Getter struct{
	jwt  JWTValidator
}

func (g *Getter) GetId(ctx context.Context,token string) (int, error){
	ctx = logger.AppendCtx(ctx, MethodName, "GetId")
	id, err := g.jwt.GetId(token)
	if err != nil{
		logger.GetFromCtx(ctx).ErrorContext(ctx, ErrJWTGetId, err)
		return -1, err
	}
	return id, nil
}


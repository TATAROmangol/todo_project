package service

import (
	"auth/pkg/logger"
	"context"
)

type JWTValidator interface{
	GetId(string) (int, error)
}

type Getter struct{
	repo Repo
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


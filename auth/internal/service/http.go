package service

import (
	"auth/pkg/logger"
	"context"
)

type JWTGenerator interface {
	GenerateToken(int) (string, error)
}

type Auth struct {
	repo Repo
	jwt  JWTGenerator
}

func (s *Auth) Register(ctx context.Context, log, pas string) (string, error) {
	ctx = logger.AppendCtx(ctx, MethodName, "Register")
	exist, err := s.repo.TakenLogin(ctx, log)
	if err != nil {
		return "", err
	}
	if exist {
		return "", err
	}

	id, err := s.repo.CreateUser(ctx, log, pas)
	if err != nil {
		return "", err
	}

	token, err := s.jwt.GenerateToken(id)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Auth) Login(ctx context.Context, log, pas string) (string, error) {
	ctx = logger.AppendCtx(ctx, MethodName, "Login")
	exist, err := s.repo.TakenLogin(ctx, log)
	if err != nil {
		return "", err
	}
	if !exist {
		return "", err
	}

	id, err := s.repo.CheckPassword(ctx, log, pas)
	if err != nil {
		return "", err
	}
	if id == -1 {
		return "", err
	}

	token, err := s.jwt.GenerateToken(id)
	if err != nil {
		logger.GetFromCtx(ctx).ErrorContext(ctx, ErrJWTGetId, err)
		return "", err
	}

	return token, nil
}

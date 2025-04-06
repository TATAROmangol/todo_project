package service

import (
	"auth/internal/errors"
	"fmt"
)

type JWTGenerator interface {
	GenerateToken(int) (string, error)
}

type Auth struct {
	repo Repo
	jwt  JWTGenerator
}

func (s *Auth) Register(log, pas string) (string, error) {
	exist, err := s.repo.TakenLogin(log)
	if err != nil {
		return "", fmt.Errorf("failed in db: %v", err)
	}
	if exist {
		return "", errors.ErrLoginTaken
	}

	id, err := s.repo.CreateUser(log, pas)
	if err != nil {
		return "", fmt.Errorf("failed in db: %v", err)
	}

	token, err := s.jwt.GenerateToken(id)
	if err != nil {
		return "", fmt.Errorf("failed in jwt: %v", err)
	}

	return token, nil
}

func (s *Auth) Login(log, pas string) (string, error) {
	exist, err := s.repo.TakenLogin(log)
	if err != nil {
		return "", fmt.Errorf("failed in db: %v", err)
	}
	if !exist {
		return "", errors.ErrUnknownLogin
	}

	id, err := s.repo.CheckPassword(log, pas)
	if err != nil {
		return "", fmt.Errorf("failed in db: %v", err)
	}
	if id == -1 {
		return "", errors.ErrIncorrectPassword
	}

	token, err := s.jwt.GenerateToken(id)
	if err != nil {
		return "", fmt.Errorf("failed in jwt: %v", err)
	}

	return token, nil
}

package service

import "context"

type Repo interface {
	TakenLogin(context.Context, string) (bool, error)
	CreateUser(context.Context, string, string) (int, error)
	CheckPassword(context.Context, string, string) (int, error)
}

type JWT interface{
	JWTGenerator
	JWTValidator
}

type Service struct{
	Auth Auth
	Getter Getter
}

func NewService(repo Repo, jwt JWT) *Service{
	return &Service{
		Auth: Auth{repo, jwt},
		Getter: Getter{repo, jwt},
	}
}
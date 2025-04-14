package service

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
		Getter: Getter{jwt},
	}
}
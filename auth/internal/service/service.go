package service

type Repo interface {
	TakenLogin(string) (bool, error)
	CreateUser(string, string) (int, error)
	CheckPassword(string, string) (int, error)
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
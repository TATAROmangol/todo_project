package service

type JWTValidator interface{
	GetId(string) (int, error)
}

type Getter struct{
	repo Repo
	jwt  JWTValidator
}

func (g *Getter) GetId(token string) (int, error){
	return g.jwt.GetId(token)
}


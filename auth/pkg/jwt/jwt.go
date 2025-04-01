package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	key []byte
}

func New() (*JWT, error) {
	key, exist := os.LookupEnv("JWT_KEY")
	if !exist {
		return nil, fmt.Errorf("failed load JWT_KEY")
	}

	return &JWT{[]byte(key)}, nil
}

func (j *JWT) GenerateToken(id int) (string, error) {
	claims := jwt.MapClaims{
		"id": id,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.key)
}

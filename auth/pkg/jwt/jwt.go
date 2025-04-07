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


func (j *JWT) GetId(tokenString string) (int, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return j.key, nil
	})

	if err != nil {
		return -1, fmt.Errorf("failed in parse token: %v", err)
	}

	if !token.Valid{
		return -1, fmt.Errorf("invalid token: %v", err)
	}

	if exp, ok := claims["exp"].(float64); ok {
        expirationTime := time.Unix(int64(exp), 0)
        if expirationTime.Before(time.Now()) {
            return -1, fmt.Errorf("token has expired")
        }
    } else {
        return -1, fmt.Errorf("token does not contain exp")
    }

	id, ok := claims["id"].(float64)
	if !ok{
		return -1, fmt.Errorf("token does not contain id")
	}
	return int(id), nil
}
package jwt

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	Key = ""
)

func MustLoad() {
	key, exist := os.LookupEnv("JWT_KEY")
	if !exist {
		log.Fatal("no found env JWT_KEY")
	}

	Key = key
}

func GetId(tokenString string) (int, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return Key, nil
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

	if _, ok := claims["id"]; !ok{
		return -1, fmt.Errorf("token does not contain id")
	}
	return claims["id"].(int), nil
}

package auth

import (
	"fmt"
	"os"
)

type Config struct {
	Address string
}

func Load() (Config, error) {
	host, exist := os.LookupEnv("NGINX_HOST")
	if !exist {
		return Config{}, fmt.Errorf("no found env NGINX_HOST")
	}
	port, exist := os.LookupEnv("NGINX_PORT")
	if !exist {
		return Config{}, fmt.Errorf("no found env NGINX_PORT")
	}
	httpAddress := fmt.Sprintf("%v:%v", host, port)

	return Config{
		Address: httpAddress,
	}, nil
}

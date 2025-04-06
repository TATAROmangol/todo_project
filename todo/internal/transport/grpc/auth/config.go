package auth

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
	Address string
}

func MustLoadConfig() Config {
	host, exist := os.LookupEnv("NGINX_HOST")
	if !exist {
		log.Fatal("no found env NGINX_HOST")
	}
	port, exist := os.LookupEnv("NGINX_PORT")
	if !exist {
		log.Fatal("no found env NGINX_PORT")
	}
	httpAddress := fmt.Sprintf("%v:%v", host, port)

	return Config{
		Address: httpAddress,
	}
}
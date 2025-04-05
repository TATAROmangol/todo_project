package v1

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
	Address string
}

func MustLoadConfig() Config {
	host, exist := os.LookupEnv("HTTP_HOST")
	if !exist {
		log.Fatal("no found env HTTP_HOST")
	}
	port, exist := os.LookupEnv("HTTP_PORT")
	if !exist {
		log.Fatal("no found env HTTP_PORT")
	}
	httpAddress := fmt.Sprintf("%v:%v", host, port)

	return Config{
		Address: httpAddress,
	}
}

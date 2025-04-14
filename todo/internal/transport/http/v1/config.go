package v1

import (
	"fmt"
	"os"
)

type Config struct {
	Address string
}

func Load() (Config, error) {
	host, exist := os.LookupEnv("HTTP_HOST")
	if !exist {
		return Config{}, fmt.Errorf("no found env HTTP_HOST")
	}
	port, exist := os.LookupEnv("HTTP_PORT")
	if !exist {
		return Config{}, fmt.Errorf("no found env HTTP_PORT")
	}
	httpAddress := fmt.Sprintf("%v:%v", host, port)

	return Config{
		Address: httpAddress,
	}, nil
}

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
	httpPort, exist := os.LookupEnv("HTTP_PORT")
	if !exist {
		log.Fatal("no found env HTTP_PORT")
	}
	httpAddress := fmt.Sprintf(":%v", httpPort)

	return Config{
		Address: httpAddress,
	}
}

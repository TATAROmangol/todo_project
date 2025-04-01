package gv1

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct{
	Host string 
	Port int
}

func LoadConfig() (Config, error){
	host, exist := os.LookupEnv("GRPC_HOST")
	if !exist{
		return Config{}, fmt.Errorf("failed to load GRPC_HOST")
	}

	sPort, exist := os.LookupEnv("GRPC_PORT")
	if !exist{
		return Config{}, fmt.Errorf("failed to load GRPC_PORT")
	}
	port, err := strconv.Atoi(sPort)
	if err != nil{
		return Config{}, fmt.Errorf("failed to load GRPC_PORT")
	}

	return Config{
		host,
		port,
	}, nil
}

func (cfg Config) GetConnectPath() string{
	return fmt.Sprintf("%v:%v", cfg.Host, cfg.Port)
}
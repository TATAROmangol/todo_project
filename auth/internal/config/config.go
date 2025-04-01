package config

import (
	"auth/internal/transport/grpc/gv1"
	"auth/pkg/postgres"
	"log"
)

type Config struct {
	GRPC gv1.Config
	PG   postgres.Config
}

func MustLoad() Config {
	grpcCfg, err := gv1.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load grpc config: %v", err)
	}

	pgCfg, err := postgres.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load postgres config: %v", err)
	}

	return Config{
		GRPC: grpcCfg,
		PG:   pgCfg,
	}
}

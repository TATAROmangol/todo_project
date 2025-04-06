package config

import (
	"auth/internal/transport/grpc/gv1"
	v1 "auth/internal/transport/http/v1"
	"auth/pkg/postgres"
	"log"
)

type Config struct {
	HTTP v1.Config
	GRPC gv1.Config
	PG   postgres.Config
}

func MustLoad() Config {
	http, err := v1.Load()
	if err != nil {
		log.Fatalf("failed to load http config: %v", err)
	}

	grpc, err := gv1.Load()
	if err != nil {
		log.Fatalf("failed to load grpc config: %v", err)
	}

	pg, err := postgres.Load()
	if err != nil {
		log.Fatalf("failed to load postgres config: %v", err)
	}

	return Config{
		HTTP: http,
		GRPC: grpc,
		PG:   pg,
	}
}

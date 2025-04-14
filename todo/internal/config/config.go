package config

import (
	"log"
	"todo/internal/transport/grpc/auth"
	v1 "todo/internal/transport/http/v1"
	"todo/pkg/migrator"
	"todo/pkg/postgres"
)

type Config struct {
	Auth auth.Config
	Http v1.Config
	Repo postgres.Config
	Migrator migrator.Config
}

func MustLoad() Config {
	authConfig, err := auth.Load()
	if err != nil {
		log.Fatalf("failed to load auth config: %v", err)
	}

	httpConfig, err := v1.Load()
	if err != nil {
		log.Fatalf("failed to load http config: %v", err)
	}

	repoConfig, err := postgres.Load()
	if err != nil {
		log.Fatalf("failed to load repo config: %v", err)
	}

	return Config{
		Auth: authConfig,
		Http: httpConfig,
		Repo: repoConfig,
	}
}

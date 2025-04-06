package config

import (
	"todo/internal/transport/grpc/auth"
	v1 "todo/internal/transport/http/v1"
	"todo/pkg/postgres"
)

type Config struct {
	Auth auth.Config
	Http v1.Config
	Repo postgres.Config
}

func MustLoad() Config {
	authConfig := auth.MustLoadConfig()
	httpConfig := v1.MustLoadConfig()
	repoConfig := postgres.MustLoadConfig()

	return Config{
		Auth: authConfig,
		Http: httpConfig,
		Repo: repoConfig,
	}
}

package config

import (
	v1 "todo/internal/transport/http/v1"
	"todo/pkg/postgres"
)

type Config struct {
	Http v1.Config
	Repo postgres.Config
}

func MustLoad() Config {
	httpConfig := v1.MustLoadConfig()
	repoConfig := postgres.MustLoadConfig()

	return Config{
		Http: httpConfig,
		Repo: repoConfig,
	}
}

package main

import (
	"auth/internal/config"
	"auth/internal/repository"
	"auth/internal/service"
	"auth/internal/transport/grpc/gv1"
	"auth/pkg/jwt"
	"auth/pkg/logger"
	"auth/pkg/postgres"
	"auth/pkg/postgres/migrator"
	"context"
	"os"
)

const (
	migrationPath = "file://internal/repository/migrations"
)

func main() {
	cfg := config.MustLoad()

	ctx := context.Background()

	l := logger.New()

	db, err := postgres.NewConnect(cfg.PG)
	if err != nil {
		l.ErrorContext(ctx, "failed to connect postgres", "err", err)
		os.Exit(1)
	}
	defer db.Close()

	m, err := migrator.New(migrationPath, cfg.PG)
	if err != nil {
		l.ErrorContext(ctx, "failed in create migrator", "error", err)
		os.Exit(1)
	}
	l.InfoContext(ctx, "migrator loaded")

	if err := m.Up(); err != nil {
		l.ErrorContext(ctx, "failed in up migrate", "error", err)
		os.Exit(1)
	}
	l.InfoContext(ctx, "migrations complete")

	repo := repository.New(db)

	jwt, err := jwt.New()
	if err != nil{
		l.ErrorContext(ctx, "failed in generate jwt", "error", err)
		os.Exit(1)
	}

	service := service.NewService(repo, jwt)
	server := gv1.New(ctx, cfg.GRPC, l, service)
	server.Run()
}

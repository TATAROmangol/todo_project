package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo/internal/config"
	"todo/internal/repository"
	service "todo/internal/services"
	"todo/internal/transport/grpc/auth"
	v1 "todo/internal/transport/http/v1"
	"todo/pkg/logger"
	"todo/pkg/migrator"
	"todo/pkg/postgres"
)

const (
	migrationPath = "file://internal/repository/migrations"
)

func main() {
	cfg := config.MustLoad()

	ctx := context.Background()

	l := logger.New()
	ctx = logger.InitFromCtx(ctx, l)

	pq, err := postgres.NewDB(cfg.Repo)
	if err != nil {
		l.ErrorContext(ctx, "failed in initialize storage", err)
		os.Exit(1)
	}
	l.InfoContext(ctx, "database loaded")

	m, err := migrator.New(migrationPath, cfg.Migrator)
	if err != nil {
		l.ErrorContext(ctx, "failed in create migrator", err)
		os.Exit(1)
	}
	l.InfoContext(ctx, "migrator loaded")

	if err := m.Up(); err != nil {
		l.ErrorContext(ctx, "failed in up migrate", err)
		os.Exit(1)
	}
	l.InfoContext(ctx, "migrations complete")

	taskRepo := repository.NewRepository(pq)
	taskService := service.NewService(taskRepo)

	auther := auth.NewAuthClient(cfg.Auth)

	router := v1.New(ctx, cfg.Http, taskService, auther)

	go func() {
		if err := router.Run(); err != nil {
			l.ErrorContext(ctx, "failed in server", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	<-c

	l.InfoContext(ctx, "started shutdown")

	closeCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	router.Shutdown(closeCtx)
	l.InfoContext(ctx, "server stop")

	pq.Close()
}

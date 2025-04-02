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
	"os/signal"
	"syscall"
	"time"
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

	go func(){
		if err := server.Run(); err != nil{
			l.ErrorContext(ctx, "failed in run server", "error", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	<- c 
	l.InfoContext(ctx, "start graceful stop")

	
	timer := time.NewTimer(10 * time.Second)

	go func(){
		<- timer.C
		l.ErrorContext(ctx, "failed in run server", "error", err)
		os.Exit(1)
	}()

	server.GracefulStop()
	db.Close()
	timer.Stop()

	l.InfoContext(ctx, "end graceful stop")
}

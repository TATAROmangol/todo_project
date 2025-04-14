package main

import (
	"auth/internal/config"
	"auth/internal/repository"
	"auth/internal/service"
	"auth/internal/transport/grpc/gv1"
	v1 "auth/internal/transport/http/v1"
	"auth/pkg/jwt"
	"auth/pkg/logger"
	"auth/pkg/postgres"
	"auth/pkg/migrator"
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
	ctx = logger.InitFromCtx(ctx, l)

	db, err := postgres.NewDB(cfg.PG)
	if err != nil {
		l.ErrorContext(ctx, "failed to connect postgres", err)
		os.Exit(1)
	}
	defer db.Close()

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

	repo := repository.New(db)

	jwt, err := jwt.New()
	if err != nil{
		l.ErrorContext(ctx, "failed in generate jwt", err)
		os.Exit(1)
	}

	service := service.NewService(repo, jwt)
	grpcServer := gv1.New(ctx, cfg.GRPC, &service.Getter)

	handler := v1.NewAuthHandler(&service.Auth)
	httpServer := v1.New(ctx, cfg.HTTP, handler)

	go func(){
		if err := grpcServer.Run(); err != nil{
			l.ErrorContext(ctx, "failed in run server", err)
			os.Exit(1)
		}
	}()

	go func(){
		if err := httpServer.Run(); err != nil{
			l.ErrorContext(ctx, "failed in run server", err)
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
		l.ErrorContext(ctx, "failed in run server", err)
		os.Exit(1)
	}()

	httpServer.Shutdown(ctx)
	grpcServer.GracefulStop()
	db.Close()
	timer.Stop()

	l.InfoContext(ctx, "end graceful stop")
}

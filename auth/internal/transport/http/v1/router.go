package v1

import (
	"context"
	"net/http"
	"auth/pkg/logger"
)

type Service interface {
	AuthService
}

type Auther interface{
	Register(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)
	LogOut(http.ResponseWriter, *http.Request)
}

type Router struct {
	ctx context.Context
	cfg Config
	srv *http.Server
}

func New(ctx context.Context, cfg Config, as Auther) *Router {
	mux := http.NewServeMux()
	mux.HandleFunc("/register", InitLoggerCtx(ctx, Operation(as.Register)))
	mux.HandleFunc("/login", InitLoggerCtx(ctx, Operation(as.Login)))
	mux.HandleFunc("/logout", InitLoggerCtx(ctx, Operation(as.LogOut)))

	srv := &http.Server{
		Addr:    cfg.Address,
		Handler: mux,
	}
	return &Router{ctx, cfg, srv}
}

func (r *Router) Run() error {
	logger.GetFromCtx(r.ctx).InfoContext(r.ctx, "Run http", "path",r.cfg.Address)
	if err := r.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (r *Router) Shutdown(ctx context.Context) error {
	if err := r.srv.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}

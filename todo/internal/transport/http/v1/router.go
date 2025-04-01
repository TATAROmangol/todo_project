package v1

import (
	"context"
	"net/http"
	"todo/pkg/logger"
)

type Service interface {
	TaskService
}

type Router struct {
	ctx context.Context
	cfg Config
	l *logger.Logger
	srv *http.Server
}

func New(ctx context.Context, cfg Config, log *logger.Logger, cases Service) *Router {
	ctx = logger.AppendCtx(ctx, "path", cfg.Address)
	
	th := NewTaskHandler(ctx, cases)

	mux := http.NewServeMux()
	mux.HandleFunc("/post", th.Post)
	mux.HandleFunc("/delete", th.Remove)
	mux.HandleFunc("/get", th.Get)

	srv := &http.Server{
		Addr: cfg.Address,
		Handler: mux,
	}
	return &Router{ctx, cfg, log, srv}
}

func (r *Router) Run() error{
	r.l.InfoContext(r.ctx, "Run server")
	if err := r.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed{
		return err
	}
	return nil
}

func (r *Router) Shutdown(ctx context.Context) error{
	if err := r.srv.Shutdown(ctx); err != nil{
		return err
	}
	return nil
}

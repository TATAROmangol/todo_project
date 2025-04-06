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
	srv *http.Server
}

func New(ctx context.Context, cfg Config, cases Service, auther Auther) *Router {
	th := NewTaskHandler(cases)

	mux := http.NewServeMux()
	mux.HandleFunc("/post", InitLoggerCtx(ctx, Operation(Auth(auther, th.Post))))
	mux.HandleFunc("/delete", InitLoggerCtx(ctx, Operation(Auth(auther, th.Remove))))
	mux.HandleFunc("/get", InitLoggerCtx(ctx, Operation(Auth(auther, th.Get))))

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

package logger

import (
	"context"
	"log/slog"
)

const(
	log_fields = "log_fields"
)

type ContextHandler struct {
    slog.Handler
}

func (h ContextHandler) Handle(ctx context.Context, r slog.Record) error {
    if attrs, ok := ctx.Value(log_fields).([]slog.Attr); ok {
        for _, v := range attrs {
            r.AddAttrs(v)
        }
    }

    return h.Handler.Handle(ctx, r)
}

func AppendCtx(ctx context.Context, key string, val any) context.Context {
	attr := slog.Any(key, val)

    if ctx == nil {
        ctx = context.Background()
    }

    if v, ok := ctx.Value(log_fields).([]slog.Attr); ok {
        v = append(v, attr)
        return context.WithValue(ctx, log_fields, v)
    }

    v := []slog.Attr{}
    v = append(v, attr)
    return context.WithValue(ctx, log_fields, v)
}
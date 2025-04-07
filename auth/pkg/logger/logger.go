package logger

import (
	"context"
	"log/slog"
	"os"
)

const(
	Key = "logger"
)

type Logger struct{
	log *slog.Logger
}

func New() *Logger{
	log := slog.New(
		&ContextHandler{slog.NewJSONHandler(os.Stdout, nil)},
	)
	return &Logger{log}
}

func InitFromCtx(ctx context.Context, log *Logger) context.Context{
	return context.WithValue(ctx, Key, log)
}

func GetFromCtx(ctx context.Context) *Logger{
	return ctx.Value(Key).(*Logger)
}

func (l Logger) InfoContext(ctx context.Context, msg string, args ...any){
	l.log.InfoContext(ctx, msg, args...)
}

func (l Logger) ErrorContext(ctx context.Context, msg string, err error){
	l.log.InfoContext(ctx, msg, "error", err)
}
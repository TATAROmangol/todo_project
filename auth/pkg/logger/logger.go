package logger

import (
	"context"
	"log/slog"
	"os"
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

func (l Logger) InfoContext(ctx context.Context, msg string, args ...any){
	l.log.InfoContext(ctx, msg, args...)
}

func (l Logger) ErrorContext(ctx context.Context, msg string, args ...any){
	l.log.InfoContext(ctx, msg, args...)
}
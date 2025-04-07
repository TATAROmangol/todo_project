package logger

import "log/slog"

func GetFromTests() *Logger{
	log := slog.New(
		&slog.TextHandler{},
	)
	return &Logger{log}
}
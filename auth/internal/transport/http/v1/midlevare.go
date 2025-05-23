package v1

import (
	"context"
	"net/http"
	"auth/pkg/logger"

	"github.com/google/uuid"
)

func InitLoggerCtx(ctx context.Context, h func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(ctx)
		h(w, r)
	}
}

func Operation(h func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		operationId := uuid.NewString()

		ctx := r.Context()
		ctx = logger.AppendCtx(ctx, OperationKey, operationId)
		ctx = logger.AppendCtx(ctx, "method path", r.URL.Path)
		logger.GetFromCtx(ctx).InfoContext(ctx, "called method")

		r = r.WithContext(ctx)
		h(w, r)
	}
}

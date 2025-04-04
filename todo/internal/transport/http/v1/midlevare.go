package v1

import (
	"context"
	"net/http"
	"todo/pkg/logger"

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

		logger.GetFromCtx(ctx).InfoContext(ctx, "called method")

		r = r.WithContext(ctx)
		h(w, r)
	}
}

func Auth(h func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// operationId := uuid.NewString()

		// ctx = logger.AppendCtx(ctx, OperationKey, operationId)
		// logger.GetFromCtx(ctx).InfoContext(ctx, "called method")

		// r = r.WithContext(ctx)

		h(w, r)
	}
}

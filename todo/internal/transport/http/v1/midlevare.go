package v1

import (
	"context"
	"net/http"
	"todo/pkg/jwt"
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
		cookie, err := r.Cookie("user_jwt")
		if err != nil{
			logger.GetFromCtx(r.Context()).ErrorContext(r.Context(), "not found cookie", "error", err)
			return
		}

		id, err := jwt.GetId(cookie.Value)
		if err != nil{
			logger.GetFromCtx(r.Context()).ErrorContext(r.Context(), "", "error", err)
			return
		}

		ctx := r.Context()
		ctx = logger.AppendCtx(ctx, UserIdKey, id)
		r = r.WithContext(ctx)

		h(w, r)
	}
}

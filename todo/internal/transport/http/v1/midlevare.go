package v1

import (
	"context"
	"net/http"
	"todo/pkg/logger"

	"github.com/google/uuid"
)

type Auther interface {
	GetId(context.Context, string) (int, error)
}

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

func Auth(auther Auther, h func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.GetFromCtx(r.Context()).InfoContext(r.Context(), "check auth")

		cookie, err := r.Cookie("jwt_id")
		if err != nil {
			WriteError(w, err, 401)
			return
		}

		id, err := auther.GetId(r.Context(), cookie.Value)
		if err != nil {
			logger.GetFromCtx(r.Context()).ErrorContext(r.Context(), "failed in auth", "error", err)
			return
		}

		ctx := r.Context()
		ctx = logger.AppendCtx(ctx, UserIdKey, id)
		ctx = context.WithValue(ctx, UserIdKey, id)
		r = r.WithContext(ctx)

		h(w, r)
	}
}

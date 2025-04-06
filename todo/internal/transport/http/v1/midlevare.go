package v1

import (
	"context"
	"net/http"
	"todo/pkg/logger"

	"github.com/google/uuid"
)

type Auther interface{
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

		logger.GetFromCtx(ctx).InfoContext(ctx, "called method")

		r = r.WithContext(ctx)
		h(w, r)
	}
}

func Auth(auther Auther, h func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("user_jwt")
		if err != nil{
			logger.GetFromCtx(r.Context()).ErrorContext(r.Context(), "not found cookie", "error", err)
			return
		}

		id, err := auther.GetId(r.Context(), cookie.Value)
		if err != nil{
			logger.GetFromCtx(r.Context()).ErrorContext(r.Context(), "failed in auth", "error", err)
			return
		}

		ctx := r.Context()
		ctx = logger.AppendCtx(ctx, UserIdKey, id)
		r = r.WithContext(ctx)

		h(w, r)
	}
}

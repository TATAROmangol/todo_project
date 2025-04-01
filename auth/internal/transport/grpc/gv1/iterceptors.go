package gv1

import (
	"auth/pkg/logger"
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

func LoggerInterceptor(pCtx context.Context, l *logger.Logger) func(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (any, error) {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (any, error) {

		operatiodID := uuid.New()
		pCtx = logger.AppendCtx(pCtx, "operation_id", operatiodID.String())
		pCtx = logger.AppendCtx(pCtx, "method", info.FullMethod)

		l.InfoContext(pCtx, "grpc server call")

		resp, err := handler(ctx, req)
		if err != nil {
			l.ErrorContext(pCtx, "error", err)
			return nil, err
		}

		return resp, nil
	}
}

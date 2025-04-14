package service

import (
	"auth/internal/service/mock"
	"auth/pkg/logger"
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetter_GetId(t *testing.T) {
	validator := mock.NewMockJWTValidator(gomock.NewController(t))

	ctx := context.Background()
	l := logger.New()
	ctx = logger.InitFromCtx(ctx, l)

	type mockBehavior func(token string)

	tests := []struct {
		name    string
		token    string
		mockBehavior mockBehavior
		want    int
		wantErr bool
	}{
		{
			name: "ok",
			token: "test",
			mockBehavior: func(token string) {
				validator.EXPECT().
					GetId(token).
					Return(1, nil)
			},
			want: 1,
			wantErr: false,
		},
		{
			name: "not ok",
			token: "test",
			mockBehavior: func(token string) {
				validator.EXPECT().
					GetId(token).
					Return(-1, fmt.Errorf("error"))
			},
			want: -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Getter{
				jwt: validator,
			}
			tt.mockBehavior(tt.token)

			got, err := g.GetId(ctx, tt.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("Getter.GetId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Getter.GetId() = %v, want %v", got, tt.want)
			}
		})
	}
}

package gv1

import (
	"auth/internal/transport/grpc/gv1/mocks"
	ssov1 "auth/pkg/grpc/auth"
	"context"
	"reflect"
	"testing"
)

func TestApi_Login(t *testing.T) {
	auth := mocks.NewAuth(t)
	
	type args struct {
		ctx context.Context
		in  *ssov1.LoginRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *ssov1.TokenResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Api{
				UnimplementedAuthServer: tt.fields.UnimplementedAuthServer,
				service:                 tt.fields.service,
			}
			got, err := s.Login(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Api.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Api.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

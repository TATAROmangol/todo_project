package service

import (
	"auth/internal/service/mock"
	"auth/pkg/logger"
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestAuth_Register(t *testing.T) {
	repo := mock.NewMockRepo(gomock.NewController(t))
	jwt := mock.NewMockJWTGenerator(gomock.NewController(t))
	ctx := context.Background()
	l := logger.New()
	ctx = logger.InitFromCtx(ctx, l)

	type repoBehavior func(log, pas string)
	type jwtBehavior func(id int)

	type args struct {
		log string
		pas string
	}
	tests := []struct {
		name         string
		repoBehavior repoBehavior
		jwtBehavior  jwtBehavior
		args         args
		want         string
		wantErr      bool
	}{
		{
			name: "ok",
			repoBehavior: func(log, pas string) {
				repo.EXPECT().
					TakenLogin(gomock.Any(), log).
					Return(false, nil)
				repo.EXPECT().
					CreateUser(gomock.Any(), log, pas).
					Return(1, nil)
			},
			jwtBehavior: func(id int) {
				jwt.EXPECT().
					GenerateToken(id).
					Return("token", nil)
			},
			args: args{
				log: "test",
				pas: "test",
			},
			want:    "token",
			wantErr: false,
		},
		{
			name: "failed taken login without error",
			repoBehavior: func(log, pas string) {
				repo.EXPECT().
					TakenLogin(gomock.Any(), log).
					Return(true, nil)
			},
			jwtBehavior: func(id int) {},
			args: args{
				log: "test",
				pas: "test",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "failed taken login with error",
			repoBehavior: func(log, pas string) {
				repo.EXPECT().
					TakenLogin(gomock.Any(), log).
					Return(false, fmt.Errorf("error"))
			},
			jwtBehavior: func(id int) {},
			args: args{
				log: "test",
				pas: "test",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "failed create user",
			repoBehavior: func(log, pas string) {
				repo.EXPECT().
					TakenLogin(gomock.Any(), log).
					Return(false, nil)
				repo.EXPECT().
					CreateUser(gomock.Any(), log, pas).
					Return(-1, fmt.Errorf("error"))
			},
			jwtBehavior: func(id int) {},
			args: args{
				log: "test",
				pas: "test",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "failed generate token",
			repoBehavior: func(log, pas string) {
				repo.EXPECT().
					TakenLogin(gomock.Any(), log).
					Return(false, nil)
				repo.EXPECT().
					CreateUser(gomock.Any(), log, pas).
					Return(1, nil)
			},
			jwtBehavior: func(id int) {
				jwt.EXPECT().
					GenerateToken(id).
					Return("", fmt.Errorf("error"))
			},
			args: args{
				log: "test",
				pas: "test",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Auth{
				repo: repo,
				jwt:  jwt,
			}
			tt.repoBehavior(tt.args.log, tt.args.pas)
			tt.jwtBehavior(1)
			got, err := s.Register(ctx, tt.args.log, tt.args.pas)
			if (err != nil) != tt.wantErr {
				t.Errorf("Auth.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Auth.Register() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuth_Login(t *testing.T) {
	repo := mock.NewMockRepo(gomock.NewController(t))
	jwt := mock.NewMockJWTGenerator(gomock.NewController(t))
	ctx := context.Background()
	l := logger.New()
	ctx = logger.InitFromCtx(ctx, l)


	type repoBehavior func(log, pas string)
	type jwtBehavior func(id int)

	type args struct {
		log string
		pas string
	}
	tests := []struct {
		name    string
		repoBehavior repoBehavior
		jwtBehavior  jwtBehavior
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "ok",
			repoBehavior: func(log, pas string) {
				repo.EXPECT().
					TakenLogin(gomock.Any(), log).
					Return(true, nil)
				repo.EXPECT().
					CheckPassword(gomock.Any(), log, pas).
					Return(1, nil)
			},
			jwtBehavior: func(id int) {
				jwt.EXPECT().
					GenerateToken(id).
					Return("token", nil)
			},
			args: args{
				log: "test",
				pas: "test",
			},
			want: "token",
			wantErr: false,
		},
		{
			name: "failed taken login without error",
			repoBehavior: func(log, pas string) {
				repo.EXPECT().
					TakenLogin(gomock.Any(), log).
					Return(false, nil)
			},
			jwtBehavior: func(id int) {},
			args: args{
				log: "test",
				pas: "test",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "failed taken login with error",
			repoBehavior: func(log, pas string) {
				repo.EXPECT().
					TakenLogin(gomock.Any(), log).
					Return(false, fmt.Errorf("error"))
			},
			jwtBehavior: func(id int) {},
			args: args{
				log: "test",
				pas: "test",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "failed create user",
			repoBehavior: func(log, pas string) {
				repo.EXPECT().
					TakenLogin(gomock.Any(), log).
					Return(true, nil)
				repo.EXPECT().
					CheckPassword(gomock.Any(), log, pas).
					Return(-1, fmt.Errorf("error"))
			},
			jwtBehavior: func(id int) {},
			args: args{
				log: "test",
				pas: "test",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "failed generate token",
			repoBehavior: func(log, pas string) {
				repo.EXPECT().
					TakenLogin(gomock.Any(), log).
					Return(true, nil)
				repo.EXPECT().
					CheckPassword(gomock.Any(), log, pas).
					Return(1, nil)
			},
			jwtBehavior: func(id int) {
				jwt.EXPECT().
					GenerateToken(id).
					Return("", fmt.Errorf("error"))
			},
			args: args{
				log: "test",
				pas: "test",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Auth{
				repo: repo,
				jwt:  jwt,
			}
			tt.repoBehavior(tt.args.log, tt.args.pas)
			tt.jwtBehavior(1)
			got, err := s.Login(ctx, tt.args.log, tt.args.pas)
			if (err != nil) != tt.wantErr {
				t.Errorf("Auth.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Auth.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

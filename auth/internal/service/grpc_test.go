package service

import (
	"auth/internal/errors"
	"auth/internal/service/mocks"
	"testing"
)

func TestService_Register(t *testing.T) {
	type args struct {
		log string
		pas string
	}
	tests := []struct {
		name    string
		a       args
		al      bool
		aj      string
		wantRes string
		wantErr error
	}{
		{
			name: "ok",
			a: args{
				log: "test",
				pas: "test",
			},
			al:      false,
			aj:      "token",
			wantRes: "token",
			wantErr: nil,
		},
		{
			name: "failed login",
			a: args{
				log: "test",
				pas: "test",
			},
			al:      true,
			aj:      "",
			wantRes: "",
			wantErr: errors.ErrLoginTaken,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepo(t)
			jwt := mocks.NewJWT(t)

			repo.On("TakenLogin", tt.a.log).Return(tt.al, nil)
			if !tt.al {
				repo.On("CreateUser", tt.a.log, tt.a.pas).Return(1, nil)
				jwt.On("GenerateToken", 1).Return(tt.aj, nil)
			}

			s := &Service{
				repo: repo,
				jwt:  jwt,
			}
			got, err := s.Register(tt.a.log, tt.a.pas)
			if err != tt.wantErr {
				t.Errorf("Service.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.wantRes {
				t.Errorf("Service.Register() = %v, want %v", got, tt.wantRes)
			}
		})
	}
}

func TestService_Login(t *testing.T) {
	type args struct {
		log string
		pas string
	}
	tests := []struct {
		name    string
		a       args
		al      bool
		ap      int
		aj      string
		wantRes string
		wantErr error
	}{
		{
			name: "ok",
			a: args{
				log: "test",
				pas: "test",
			},
			al:      true,
			ap:      1,
			aj:      "token",
			wantRes: "token",
			wantErr: nil,
		},
		{
			name: "failed login",
			a: args{
				log: "test",
				pas: "test",
			},
			al:      false,
			ap:      1,
			aj:      "",
			wantRes: "",
			wantErr: errors.ErrUnknownLogin,
		},
		{
			name: "failed password",
			a: args{
				log: "test",
				pas: "test",
			},
			al:      true,
			ap:      -1,
			aj:      "",
			wantRes: "",
			wantErr: errors.ErrIncorrectPassword,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepo(t)
			jwt := mocks.NewJWT(t)

			repo.On("TakenLogin", tt.a.log).Return(tt.al, nil)
			if tt.al {
				repo.On("CheckPassword", tt.a.log, tt.a.pas).Return(tt.ap, nil)
				if tt.ap != -1 {
					jwt.On("GenerateToken", tt.ap).Return(tt.aj, nil)
				}
			}

			s := &Service{
				repo: repo,
				jwt:  jwt,
			}
			got, err := s.Login(tt.a.log, tt.a.pas)
			if err != tt.wantErr {
				t.Errorf("Service.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.wantRes {
				t.Errorf("Service.Register() = %v, want %v", got, tt.wantRes)
			}
		})
	}
}

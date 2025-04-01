package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestRepository_TakenLogin(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	r := New(db)

	type mockBehavior func(login string)

	tests := []struct {
		name         string
		login        string
		mockBehavior mockBehavior
		want         bool
		wantErr      bool
	}{
		{
			name:  "ok",
			login: "test",
			mockBehavior: func(login string) {
				mock.ExpectPrepare("SELECT EXIST").
					ExpectQuery().
					WithArgs(login).
					WillReturnRows(sqlmock.NewRows([]string{"exist"}).AddRow(true))
			},
			want:    true,
			wantErr: false,
		},
		{
			name:  "not have",
			login: "test",
			mockBehavior: func(login string) {
				mock.ExpectPrepare("SELECT EXIST").
					ExpectQuery().
					WithArgs(login).
					WillReturnRows(sqlmock.NewRows([]string{"exist"}).AddRow(false))
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(tt.login)
			got, err := r.TakenLogin(tt.login)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.TakenLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Repository.TakenLogin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	r := New(db)

	type args struct {
		login    string
		password string
	}

	type mockBehavior func(args args)

	tests := []struct {
		name         string
		args         args
		mockBehavior mockBehavior
		want         int
		wantErr      error
	}{
		{
			name: "ok",
			args: args{
				"test",
				"test",
			},
			mockBehavior: func(args args) {
				mock.ExpectPrepare("INSERT INTO users").
					ExpectQuery().
					WithArgs(args.login, args.password).
					WillReturnRows(sqlmock.NewRows([]string{"exist"}).AddRow(1))
			},
			want:    1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(tt.args)
			got, err := r.CreateUser(tt.args.login, tt.args.password)
			if err != tt.wantErr {
				t.Errorf("Repository.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Repository.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_CheckPassword(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	r := New(db)

	type args struct {
		login    string
		password string
	}

	type mockBehavior func(args args)

	tests := []struct {
		name         string
		args         args
		mockBehavior mockBehavior
		want         int
		wantErr      error
	}{
		{
			name: "ok",
			args: args{
				"test",
				"test",
			},
			mockBehavior: func(args args) {
				mock.ExpectPrepare("SELECT id").
					ExpectQuery().
					WithArgs(args.login, args.password).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "failed password",
			args: args{
				"test",
				"test",
			},
			mockBehavior: func(args args) {
				mock.ExpectPrepare("SELECT id").
					ExpectQuery().
					WithArgs(args.login, args.password).
					WillReturnRows(sqlmock.NewRows([]string{"id"}))
			},
			want:    -1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(tt.args)

			got, err := r.CheckPassword(tt.args.login, tt.args.password)
			if err != tt.wantErr {
				t.Errorf("Repository.CheckPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Repository.CheckPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

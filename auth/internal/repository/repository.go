package repository

import (
	"auth/pkg/logger"
	"context"
	"database/sql"
)

type Repository struct{
	db *sql.DB
}

func New(db *sql.DB) *Repository{
	return &Repository{db}
}

func (r *Repository) TakenLogin(ctx context.Context,login string) (bool, error){
	stmt, err := r.db.Prepare(`
		SELECT EXISTS(
			SELECT 1
			FROM users
			WHERE login = ($1)
		)
	`)
	if err != nil{
		logger.GetFromCtx(ctx).ErrorContext(ctx, ErrCreateSTMT, err)
		return false, err
	}

	var exist bool 
	if err := stmt.QueryRow(login).Scan(&exist); err != nil{
		logger.GetFromCtx(ctx).ErrorContext(ctx, ErrCheckLogin, err)
		return false, err
	}

	return exist, nil
}

func (r *Repository) CreateUser(ctx context.Context, login, password string) (int, error){
	stmt, err := r.db.Prepare(`
		INSERT INTO users (login, password)
		VALUES ($1, $2) 
		RETURNING id
	`)
	if err != nil {
		logger.GetFromCtx(ctx).ErrorContext(ctx, ErrCreateSTMT, err)
		return -1, err
	}
	defer stmt.Close()

	var id int
	if err := stmt.QueryRow(login, password).Scan(&id); err != nil || id == -1{
		logger.GetFromCtx(ctx).ErrorContext(ctx, ErrCreateUser, err)
		return -1, err
	}

	return id, nil
}

func (r *Repository) CheckPassword(ctx context.Context, login, password string) (int, error){
	stmt, err := r.db.Prepare(`
		SELECT id
		FROM users
		WHERE login = ($1) AND password = ($2)
	`)
	if err != nil{
		logger.GetFromCtx(ctx).ErrorContext(ctx, ErrCreateSTMT, err)
		return -1, err
	}

	id := -1
	if err := stmt.QueryRow(login, password).Scan(&id); err != nil && err != sql.ErrNoRows{
		logger.GetFromCtx(ctx).ErrorContext(ctx, ErrCheckUser, err)
		return -1, err
	}

	return id, nil
}
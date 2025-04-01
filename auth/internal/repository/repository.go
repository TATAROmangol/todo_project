package repository

import (
	"database/sql"
	"fmt"
)

type Repository struct{
	db *sql.DB
}

func New(db *sql.DB) *Repository{
	return &Repository{db}
}

func (r *Repository) TakenLogin(login string) (bool, error){
	stmt, err := r.db.Prepare(`
		SELECT EXISTS(
			SELECT 1
			FROM users
			WHERE login = ($1)
		)
	`)
	if err != nil{
		return true, fmt.Errorf("failed in check login: %v", err)
	}

	var exist bool 
	if err := stmt.QueryRow(login).Scan(&exist); err != nil{
		return true, fmt.Errorf("failed in check login: %v", err)
	}

	return exist, nil
}

func (r *Repository) CreateUser(login, password string) (int, error){
	stmt, err := r.db.Prepare(`
		INSERT INTO users (login, password)
		VALUES ($1, $2) 
		RETURNING id
	`)
	if err != nil {
		return -1, fmt.Errorf("failed create user: %v", err)
	}
	defer stmt.Close()

	var id int
	if err := stmt.QueryRow(login, password).Scan(&id); err != nil || id == -1{
		return -1, fmt.Errorf("failed create user: %v", err)
	}

	return id, nil
}

func (r *Repository) CheckPassword(login, password string) (int, error){
	stmt, err := r.db.Prepare(`
		SELECT id
		FROM users
		WHERE login = ($1) AND password = ($2)
	`)
	if err != nil{
		return -1, fmt.Errorf("failed in check user: %v", err)
	}

	id := -1
	if err := stmt.QueryRow(login, password).Scan(&id); err != nil && err != sql.ErrNoRows{
		return -1, fmt.Errorf("failed in check user: %v", err)
	}

	return id, nil
}
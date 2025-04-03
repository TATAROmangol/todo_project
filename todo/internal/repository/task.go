package repository

import (
	"context"
	"database/sql"
	"todo/internal/entities"

	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	ctx context.Context
	db  *sql.DB
}

func NewRepository(ctx context.Context, db *sql.DB) *Repository {
	return &Repository{ctx, db}
}

func (r *Repository) Get(user_id int) ([]entities.Task, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, name 
		  FROM tasks
		WHERE user_id = $1
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var res []entities.Task
	rows, err := stmt.Query(user_id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task entities.Task
		rows.Scan(&task.Id, &task.Name)
		res = append(res, task)
	}

	return res, nil
}

func (r *Repository) Create(name string, user_id int) (entities.Task, error) {
	stmt, err := r.db.Prepare(`
		INSERT INTO tasks(name)
		VALUES ($1, $2) RETURNING id
	`)
	if err != nil {
		return entities.Task{}, err
	}
	defer stmt.Close()

	var id int
	if err := stmt.QueryRow(name, user_id).Scan(&id); err != nil {
		return entities.Task{}, err
	}

	return entities.Task{Id: id, Name: name}, nil
}

func (r *Repository) Remove(id, user_id int) error {
	stmt, err := r.db.Prepare(`
		DELETE FROM tasks 
		WHERE id = $1 AND user_id=$2
		`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id, user_id); err != nil {
		return err
	}

	return nil
}

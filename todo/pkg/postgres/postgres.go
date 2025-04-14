package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewDB(cfg Config) (*sql.DB, error) {
	conStr := cfg.GetConnString()
	db, err := sql.Open("postgres", conStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect from postgresql: %v", err)
	}

	return db, nil
}
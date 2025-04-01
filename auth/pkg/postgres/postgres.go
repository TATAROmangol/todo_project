package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewConnect(cfg Config) (*sql.DB, error){
	db, err := sql.Open("postgres", cfg.GetConnectPath())
	if err != nil{
		return nil, fmt.Errorf("failed to connect postgres: %v", err)
	}

	return db, nil
}
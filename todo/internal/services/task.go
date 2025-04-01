package service

import (
	"context"
	"todo/internal/entities"
)

type Repo interface {
	Get() ([]entities.Task, error)
	Create(name string) (entities.Task, error)
	Remove(id int) error
}

type Service struct {
	ctx context.Context
	db Repo
}

func NewService(ctx context.Context, db Repo) *Service {
	return &Service{ctx, db}
}

func (tc *Service) GetTasks() ([]entities.Task, error) {
	return tc.db.Get()
}

func (tc *Service) CreateTask(name string) (entities.Task, error) {
	return tc.db.Create(name)
}

func (tc *Service) RemoveTask(id int) error {
	return tc.db.Remove(id)
}

package service

import (
	"context"
	"todo/internal/entities"
)

type Repo interface {
	Get(ctx context.Context, userId int) ([]entities.Task, error)
	Create(ctx context.Context, name string, userId int) (entities.Task, error)
	Remove(ctx context.Context, id, userId int) error
}

type Service struct {
	db Repo
}

func NewService(db Repo) *Service {
	return &Service{db}
}

func (tc *Service) GetTasks(ctx context.Context, userId int) ([]entities.Task, error) {
	return tc.db.Get(ctx, userId)
}

func (tc *Service) CreateTask(ctx context.Context, name string, userId int) (entities.Task, error) {
	return tc.db.Create(ctx, name, userId)
}

func (tc *Service) RemoveTask(ctx context.Context, id, userId int) error {
	return tc.db.Remove(ctx, id, userId)
}

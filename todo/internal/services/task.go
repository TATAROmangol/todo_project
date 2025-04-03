package service

import (
	"context"
	"todo/internal/entities"
)

type Repo interface {
	Get(int) ([]entities.Task, error)
	Create(string, int) (entities.Task, error)
	Remove(int, int) error
}

type Service struct {
	ctx context.Context
	db Repo
}

func NewService(ctx context.Context, db Repo) *Service {
	return &Service{ctx, db}
}

func (tc *Service) GetTasks(userId int) ([]entities.Task, error) {
	return tc.db.Get(userId)
}

func (tc *Service) CreateTask(name string, userId int) (entities.Task, error) {
	return tc.db.Create(name, userId)
}

func (tc *Service) RemoveTask(id, userId int) error {
	return tc.db.Remove(id, userId)
}

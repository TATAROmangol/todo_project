package v1

import (
	"context"
	"encoding/json"
	"net/http"
	"todo/internal/entities"
)

type TaskService interface {
	GetTasks() ([]entities.Task, error)
	CreateTask(name string) (entities.Task, error)
	RemoveTask(id int) error
}

type TaskHandler struct {
	ctx context.Context
	ts  TaskService
}

func NewTaskHandler(ctx context.Context, ts TaskService) *TaskHandler {
	return &TaskHandler{ctx, ts}
}

type CreateReq struct {
	Name string `json:"name"`
}

func (th *TaskHandler) Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		WriteError(w, nil, http.StatusMethodNotAllowed)
	}

	var req CreateReq
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if err != nil {
		WriteError(w, err, 404)
		return
	}

	task, err := th.ts.CreateTask(req.Name)
	if err != nil {
		WriteError(w, err, 404)
		return
	}

	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		WriteError(w, err, 404)
		return
	}
}

func (th *TaskHandler) Remove(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		WriteError(w, nil, http.StatusMethodNotAllowed)
	}

	var task entities.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		WriteError(w, err, 404)
		return
	}
	defer r.Body.Close()

	if err := th.ts.RemoveTask(task.Id); err != nil {
		WriteError(w, err, 404)
		return
	}
}

func (th *TaskHandler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		WriteError(w, nil, http.StatusMethodNotAllowed)
	}

	tasks, err := th.ts.GetTasks()
	if err != nil {
		WriteError(w, err, 404)
		return
	}

	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		WriteError(w, err, 404)
		return
	}
}

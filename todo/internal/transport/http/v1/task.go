package v1

import (
	"context"
	"encoding/json"
	"net/http"
	"todo/internal/entities"
)

type TaskService interface {
	GetTasks(ctx context.Context, userId int) ([]entities.Task, error)
	CreateTask(ctx context.Context, name string, userId int) (entities.Task, error)
	RemoveTask(ctx context.Context, id, userId int) error
}

type TaskHandler struct {
	ts TaskService
}

func NewTaskHandler(ts TaskService) *TaskHandler {
	return &TaskHandler{ts}
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

	userId := r.Context().Value(UserIdKey).(int)

	task, err := th.ts.CreateTask(r.Context(), req.Name, userId)
	if err != nil {
		WriteError(w, err, 404)
		return
	}

	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		WriteError(w, err, 404)
		return
	}

	w.WriteHeader(http.StatusOK)
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

	userId := r.Context().Value(UserIdKey).(int)

	if err := th.ts.RemoveTask(r.Context(), task.Id, userId); err != nil {
		WriteError(w, err, 404)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (th *TaskHandler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		WriteError(w, nil, http.StatusMethodNotAllowed)
	}

	userId := r.Context().Value(UserIdKey).(int)

	tasks, err := th.ts.GetTasks(r.Context(), userId)
	if err != nil {
		WriteError(w, err, 404)
		return
	}

	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		WriteError(w, err, 404)
		return
	}

	w.WriteHeader(http.StatusOK)
}

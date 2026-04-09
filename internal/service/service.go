package service

import (
	"fmt"
	"task-manager/internal/store"
	"task-manager/internal/task"
	"time"
)

type TaskService struct {
	store store.Store
}

func New(s store.Store) *TaskService {
	return &TaskService{store: s}
}

// Add validates the title, checks for duplicates, then persists
func (svc *TaskService) Add(title string) (task.Task, error) {
	err := task.ValidateTitle(title)
	if err != nil {
		return task.Task{}, fmt.Errorf("Invalid title: %w", err)
	}
	t := task.Task{
		TaskID:    0,
		Title:     title,
		Status:    "pending",
		CreatedAt: time.Now(),
	}
	task, err := svc.store.Add(t)
	if err != nil {
		return task, fmt.Errorf("task add: %w", err)
	}
	return task, nil
}

func (svc *TaskService) List(filterStatus string) ([]task.Task, error) {

	if filterStatus != "" {
		if err := task.ValidateStatus(filterStatus); err != nil {
			return nil, err
		}
	}

	tasks, err := svc.store.List(filterStatus)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (svc *TaskService) UpdateStatus(id int, status string) (task.Task, error) {
	return svc.store.UpdateStatus(id, status)
}

func (svc *TaskService) Delete(id int) error {
	return svc.store.DeleteTask(id)
}

package store

import "task-manager/internal/task"

type Store interface {
	Add(t task.Task) (task.Task, error)
	List(filterStatus string) ([]task.Task, error)
	UpdateStatus(id int, newStatus string) (task.Task, error)
	DeleteTask(id int) error
}

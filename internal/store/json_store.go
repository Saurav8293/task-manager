package store

import (
	"encoding/json"
	"fmt"
	"os"
	"task-manager/internal/task"
)

type JSONStore struct {
	filePath string
}

func NewJSONStore(filePath string) *JSONStore {
	return &JSONStore{filePath: filePath}
}

func (s *JSONStore) load() ([]task.Task, error) {
	data, err := os.ReadFile(s.filePath)
	if os.IsNotExist(err) {
		return []task.Task{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("reading task file: %w", err)
	}

	var tasks []task.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("parsing task file: %w", err)
	}
	return tasks, nil
}

func (s *JSONStore) Save(t []task.Task) error {
	data, err := json.MarshalIndent(t, "", "  ")

	if err != nil {
		return fmt.Errorf("serializing tasks: %w", err)
	}

	if err := os.WriteFile(s.filePath, data, 0644); err != nil {
		return fmt.Errorf("writing task file: %w", err)
	}

	return nil
}

func (s *JSONStore) Add(t task.Task) (task.Task, error) {
	tasks, err := s.load()

	if err != nil {
		return task.Task{}, err
	}

	// Auto-Increment ID
	maxID := 0

	for _, existing := range tasks {
		if existing.TaskID > maxID {
			maxID = existing.TaskID
		}
	}
	t.TaskID = maxID + 1

	tasks = append(tasks, t)

	if err := s.Save(tasks); err != nil {
		return task.Task{}, err
	}
	return t, nil

}

func (s *JSONStore) UpdateStatus(id int, status string) (task.Task, error) {
	tasks, err := s.load()
	if err != nil {
		return task.Task{}, err
	}

	for i, t := range tasks {
		if t.TaskID == id {
			tasks[i].Status = task.Status(status)
			if err := s.Save(tasks); err != nil {
				return task.Task{}, err
			}
			return tasks[i], nil
		}
	}
	return task.Task{}, &task.NotFoundError{ID: id}

}

func (s *JSONStore) List(filterStatus string) ([]task.Task, error) {
	tasks, err := s.load()

	if err != nil {
		return nil, err
	}

	if filterStatus == "" {
		return tasks, nil
	}

	var filtered []task.Task

	for _, t := range tasks {
		if t.Status == task.Status(filterStatus) {
			filtered = append(filtered, t)
		}
	}
	return filtered, nil
}

func (s *JSONStore) DeleteTask(taskId int) error {
	tasks, err := s.load()
	if err != nil {
		return fmt.Errorf("Loading tasks: %w", err)
	}
	for i, task := range tasks {
		if task.TaskID == taskId {
			tasks[i] = tasks[len(tasks)-1]
			tasks = tasks[:len(tasks)-1]
			return s.Save(tasks)
		}
	}
	return &task.NotFoundError{ID: taskId}
}

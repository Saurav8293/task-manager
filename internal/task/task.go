package task

import "time"

type Task struct {
	TaskID    int
	Title     string
	Status    string
	CreatedAt time.Time
}

func (t Task) CreateTask() {
}

func (t Task) ReadTask() {

}

func (t Task) UpdateTask() {

}

func (t Task) DeleteTask() {

}

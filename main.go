package main

import (
	"task-manager/internal/cli"
	"task-manager/internal/task"
)

func main() {
	store := task.NewStore("tasks.json")
	cli.Run(store)
}

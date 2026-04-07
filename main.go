package main

import (
	"task-manager/internal/cli"
	"task-manager/internal/service"
	"task-manager/internal/store"
)

func main() {
	s := store.NewJSONStore("tasks.json")
	service := service.New(s)
	cli.Run(service)
}

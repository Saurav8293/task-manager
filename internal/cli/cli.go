package cli

import (
	"fmt"
	"os"
	"task-manager/internal/service"
)

func Run(s *service.TaskService) {
	h := NewHandlers(s)

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	args := os.Args[2:]

	switch os.Args[1] {
	case "add":
		h.Add(args)

	case "list":
		h.List(args)

	case "update":
		h.Update(args)

	case "delete":
		h.Delete(args)

	default:
		fmt.Fprintf(os.Stderr, "unknown command: %q\n\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`task - a simple CLI task manager
usage:
	task add -title "buy groceries"
	task list
	`)
}

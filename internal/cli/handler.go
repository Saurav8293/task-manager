package cli

import (
	"flag"
	"fmt"
	"os"
	"task-manager/internal/service"
)

type Handlers struct {
	svc *service.TaskService
}

func NewHandlers(svc *service.TaskService) *Handlers {
	return &Handlers{svc: svc}
}

func (h *Handlers) Add(args []string) {
	fs := flag.NewFlagSet("add", flag.ExitOnError)
	title := fs.String("title", "", "task title (required)")
	fs.Parse(args)

	if *title == "" {
		fmt.Fprintln(os.Stderr, "error: - title is required")
		os.Exit(1)
	}
	t, err := h.svc.Add(*title)
	if err != nil {
		fmt.Errorf("Adding the title %w", err)
		return
	}
	fmt.Printf("added: %s\n", t)

}

func (h *Handlers) List(args []string) {
	fs := flag.NewFlagSet("list", flag.ExitOnError)
	status := fs.String("status", "", "filter: pending, in_progress, done")
	fs.Parse(args)

	tasks, _ := h.svc.List(*status)
	fmt.Println(tasks)

}

func (h *Handlers) Update(args []string) {
	fs := flag.NewFlagSet("update", flag.ExitOnError)
	id := fs.Int("id", 0, "task ID (required)")
	status := fs.String("status", "", "filter: pending, in_progress, done")
	fs.Parse(args)

	
	if *id == 0 || *status == "" {
		fmt.Fprintln(os.Stderr, "error: -id and -status required")
		os.Exit(1)
	}

	t, err := h.svc.UpdateStatus(*id, *status)
	if err != nil {
		fmt.Errorf("Updating status: %w", err)
	}
	fmt.Println("Task update successfully", t)
}

func (h *Handlers) Delete(args []string) {
	fs := flag.NewFlagSet("delete", flag.ExitOnError)
	id := fs.Int("id", 0, "task ID (required)")
	fs.Parse(args)

	if *id == 0 {
		fmt.Fprintln(os.Stderr, "err: -id required")
		os.Exit(1)
	}

	err := h.svc.Delete(*id)
	if err != nil {
		fmt.Errorf("Task Delete: %w", err)
	}
	fmt.Printf("Task deleted successfully")

}

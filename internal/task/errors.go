package task

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed on '%s': %s", e.Field, e.Message)
}

type NotFoundError struct {
	ID int
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("task with id %d not found", e.ID)
}



var (
	ErrInvalidStatus = errors.New("Invalid status")
)

func ValidateStatus(s string) error {
	switch Status(s) {
	case StatusPending, StatusInProgress, StatusDone:
		fmt.Println("status is: ", StatusDone, StatusInProgress, StatusPending)
		return nil

	default:
		return fmt.Errorf("'%s' is not valid: %w", s, ErrInvalidStatus)
	}
}

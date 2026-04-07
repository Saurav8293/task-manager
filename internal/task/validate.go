package task

func ValidateTitle(title string) error {
	if title == "" {
		return &ValidationError{Field: "title", Message: "cannot be empty"}
	}

	if len(title) > 100 {
		return &ValidationError{
			Field:   "title",
			Message: "must be 100 characters or fewer",
		}
	}
	return nil

}

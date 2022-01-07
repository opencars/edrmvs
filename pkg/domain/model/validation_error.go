package model

import (
	"fmt"
	"strings"
)

type ValidationError struct {
	Messages map[string][]string
}

func NewValidationError(messages map[string][]string) ValidationError {
	return ValidationError{
		Messages: messages,
	}
}

func (e *ValidationError) Append(field string, message ...string) {
	if e.Messages == nil {
		e.Messages = make(map[string][]string)
	}

	e.Messages[field] = append(e.Messages[field], message...)
}

func (e ValidationError) WithPrefix(prefix string) ValidationError {
	messages := make(map[string][]string)

	for k, v := range e.Messages {
		key := fmt.Sprintf("%s.%s", prefix, k)

		messages[key] = append(messages[key], v...)
	}

	return ValidationError{
		Messages: messages,
	}
}

func (e ValidationError) Error() string {
	errs := make([]string, 0)

	for k, items := range e.Messages {
		errs = append(errs, fmt.Sprintf("%s: (%s)", k, strings.Join(items, ", ")))
	}

	return strings.Join(errs, ", ")
}

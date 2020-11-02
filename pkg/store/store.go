package store

import (
	"context"
	"errors"
)

var (
	// ErrRecordNotFound returned, when entity does not exist.
	ErrRecordNotFound = errors.New("record not found")
)

// Store is an aggregation of repositories.
type Store interface {
	Registration() RegistrationRepository
	Health(ctx context.Context) error
}

package store

import (
	"errors"
)

var (
	// ErrRecordNotFound returned, when entity does not exist.
	ErrRecordNotFound = errors.New("record not found")
)

// Store is an aggregation of repositories.
type Store interface {
	Registration() RegistrationRepository
}

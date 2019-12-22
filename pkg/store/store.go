package store

import "errors"

var (
	RecordNotFound = errors.New("record not found")
)

type Store interface {
	Registration() RegistrationRepository
}

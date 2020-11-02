package teststore

import (
	"context"
	"github.com/opencars/edrmvs/pkg/store"
)

// Store is postgres wrapper for store.Store.
type Store struct {
	registrationRepository *RegistrationRepository
}

// New returns store instance.
func New() store.Store {
	return &Store{}
}

// Registration is responsible for registrations manipulation.
func (s *Store) Registration() store.RegistrationRepository {
	if s.registrationRepository == nil {
		s.registrationRepository = &RegistrationRepository{}
	}

	return s.registrationRepository
}


func (s *Store) Health(_ context.Context) error {
	return nil
}

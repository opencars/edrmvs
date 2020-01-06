package teststore

import (
	"github.com/opencars/edrmvs/pkg/store"
)

// Store is postgres wrapper for store.Store.
type Store struct {
	registrationRepository *RegistrationRepository
}

// Registration is responsible for registrations manipulation.
func (s *Store) Registration() store.RegistrationRepository {
	if s.registrationRepository == nil {
		s.registrationRepository = &RegistrationRepository{}
	}

	return s.registrationRepository
}

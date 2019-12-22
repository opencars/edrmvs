package teststore

import (
	"github.com/opencars/edrmvs/pkg/store"
)

type Store struct {
	registrationRepository *RegistrationRepository
}

func (s *Store) Registration() store.RegistrationRepository {
	if s.registrationRepository == nil {
		s.registrationRepository = &RegistrationRepository{}
	}

	return s.registrationRepository
}

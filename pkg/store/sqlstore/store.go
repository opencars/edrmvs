package sqlstore

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/opencars/edrmvs/pkg/config"
	"github.com/opencars/edrmvs/pkg/store"
)

// Store is postgres wrapper for store.Store.
type Store struct {
	db *sqlx.DB

	registrationRepository *RegistrationRepository
}

// Registration is responsible for registrations manipulation.
func (s *Store) Registration() store.RegistrationRepository {
	if s.registrationRepository == nil {
		s.registrationRepository = &RegistrationRepository{
			store: s,
		}
	}

	return s.registrationRepository
}

// New returns new instance of store.
func New(conf *config.Database) (*Store, error) {
	info := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
		conf.Host, conf.Port, conf.Username, conf.Database, conf.SSLMode, conf.Password,
	)

	db, err := sqlx.Connect("postgres", info)
	if err != nil {
		return nil, fmt.Errorf("connection failed: %w", err)
	}

	return &Store{
		db: db,
	}, nil
}

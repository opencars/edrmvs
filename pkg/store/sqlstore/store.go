package sqlstore

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/opencars/edrmvs/pkg/config"
	"github.com/opencars/edrmvs/pkg/store"
)

type Store struct {
	db                     *sqlx.DB
	registrationRepository *RegistrationRepository
}

func (s *Store) Registration() store.RegistrationRepository {
	if s.registrationRepository == nil {
		s.registrationRepository = &RegistrationRepository{
			store: s,
		}
	}

	return s.registrationRepository
}

func New(conf *config.Settings) (*Store, error) {
	var info string
	if conf.DB.Password == "" {
		info = fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
			conf.DB.Host, conf.DB.Port, conf.DB.Username, conf.DB.Database,
		)
	} else {
		info = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			conf.DB.Host, conf.DB.Port, conf.DB.Username, conf.DB.Password, conf.DB.Database,
		)
	}

	db, err := sqlx.Connect("postgres", info)
	if err != nil {
		return nil, fmt.Errorf("connection failed: %w", err)
	}

	return &Store{
		db: db,
	}, nil
}

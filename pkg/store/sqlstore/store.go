package sqlstore

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/opencars/edrmvs/pkg/config"
)

// Store is postgres wrapper for store.Store.
type RegistrationStore struct {
	db *sqlx.DB
}

func (s *RegistrationStore) Health(ctx context.Context) error {
	_, err := s.db.ExecContext(ctx, `SELECT 1`)
	if err != nil {
		return err
	}

	return nil
}

// New returns new instance of store.
func New(conf *config.Database) (*RegistrationStore, error) {
	info := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
		conf.Host, conf.Port, conf.Username, conf.Database, conf.SSLMode, conf.Password,
	)

	db, err := sqlx.Connect("postgres", info)
	if err != nil {
		return nil, fmt.Errorf("connection failed: %w", err)
	}

	db.SetConnMaxIdleTime(time.Minute)
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)

	return &RegistrationStore{
		db: db,
	}, nil
}

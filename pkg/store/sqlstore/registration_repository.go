package sqlstore

import (
	"context"
	"database/sql"
	"errors"

	"github.com/opencars/edrmvs/pkg/domain"
)

// Create adds new record to the database.
func (s *RegistrationStore) Create(ctx context.Context, registration *domain.Registration) error {
	record := convertFromDomain(registration)

	_, err := s.db.NamedExecContext(ctx,
		`INSERT INTO registrations (
			brand, capacity, color, d_first_reg, d_reg, fuel, kind,
			make_year, model, n_doc, s_doc, n_reg_new, n_seating,
			n_standing, own_weight, rank_category, total_weight, vin
		) VALUES (
			:brand, :capacity, :color, :d_first_reg, :d_reg, :fuel, :kind,
		 	:make_year, :model, :n_doc, :s_doc, :n_reg_new, :n_seating,
 		 	:n_standing, :own_weight, :rank_category, :total_weight, :vin
		) ON CONFLICT DO NOTHING`,
		record,
	)

	if err != nil {
		return err
	}

	return nil
}

// FindByNumber returns registrations with specified number.
func (s *RegistrationStore) FindByNumber(ctx context.Context, number string) ([]domain.Registration, error) {
	records := make([]Registration, 0)

	err := s.db.SelectContext(ctx, &records,
		`SELECT brand, capacity, color, d_first_reg, d_reg, fuel, kind,
			    make_year, model, CONCAT(s_doc, n_doc) as code, n_reg_new, n_seating,
			    n_standing, own_weight, rank_category, total_weight, vin
		FROM registrations
		WHERE n_reg_new = $1
		ORDER BY d_reg DESC`,
		number,
	)

	if err != nil {
		return nil, err
	}

	result := make([]domain.Registration, 0)
	for i := range records {
		result = append(result, *convertToDomain(&records[i]))
	}

	return result, nil
}

// FindByCode returns registrations with specified code.
func (s *RegistrationStore) FindByCode(ctx context.Context, code string) (*domain.Registration, error) {
	var record Registration

	err := s.db.GetContext(ctx, &record,
		`SELECT brand, capacity, color, d_first_reg, d_reg, fuel, kind,
			    make_year, model, CONCAT(s_doc, n_doc) as code, n_reg_new, n_seating,
			    n_standing, own_weight, rank_category, total_weight, vin
		FROM registrations
		WHERE CONCAT(s_doc, n_doc) = $1
		ORDER BY d_reg DESC`,
		code,
	)

	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return convertToDomain(&record), nil
}

// FindByVIN returns registrations with specified VIN.
func (s *RegistrationStore) FindByVIN(ctx context.Context, vin string) ([]domain.Registration, error) {
	records := make([]Registration, 0)

	err := s.db.SelectContext(ctx, &records,
		`SELECT brand, capacity, color, d_first_reg, d_reg, fuel, kind,
                make_year, model, CONCAT(s_doc, n_doc) as code, n_reg_new, n_seating,
                n_standing, own_weight, rank_category, total_weight, vin
		FROM registrations
		WHERE vin = $1
		ORDER BY d_reg DESC`,
		vin,
	)

	if err != nil {
		return nil, err
	}

	result := make([]domain.Registration, 0)
	for i := range records {
		result = append(result, *convertToDomain(&records[i]))
	}

	return result, nil
}

// GetLast returns last registration from the database.
func (s *RegistrationStore) FindLastBySeries(ctx context.Context, series string) (*domain.Registration, error) {
	var record Registration

	err := s.db.GetContext(ctx, &record,
		`SELECT
			brand, capacity, color, d_first_reg, d_reg, fuel,
			kind, make_year, model, s_doc, n_doc, CONCAT(s_doc, n_doc) as code, n_reg_new, n_seating,
			n_standing, own_weight, rank_category, total_weight, vin
		FROM registrations
		WHERE s_doc = $1
		ORDER BY s_doc, n_doc DESC
		LIMIT 1`,
		series,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, domain.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return convertToDomain(&record), nil
}

// AllSeries returns list of all known series.
func (s *RegistrationStore) AllSeries(ctx context.Context) ([]string, error) {
	codes := make([]string, 0)

	err := s.db.SelectContext(ctx, &codes, `SELECT s_doc FROM registrations GROUP BY s_doc`)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, domain.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return codes, nil
}

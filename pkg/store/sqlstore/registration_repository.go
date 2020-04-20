package sqlstore

import (
	"database/sql"

	"github.com/opencars/edrmvs/pkg/model"
	"github.com/opencars/edrmvs/pkg/store"
)

// RegistrationRepository is responsible for registrations manipulation.
type RegistrationRepository struct {
	store *Store
}

// Create adds new record to the database.
func (r *RegistrationRepository) Create(registration *model.Registration) error {
	_, err := r.store.db.NamedExec(
		`INSERT INTO registrations (
			brand, capacity, color, d_first_reg, d_reg, fuel, kind,
			make_year, model, n_doc, s_doc, n_reg_new, n_seating,
			n_standing, own_weight, rank_category, total_weight, vin
		) VALUES (
			:brand, :capacity, :color, :d_first_reg, :d_reg, :fuel, :kind,
		 	:make_year, :model, :n_doc, :s_doc, :n_reg_new, :n_seating,
 		 	:n_standing, :own_weight, :rank_category, :total_weight, :vin
		) ON CONFLICT DO NOTHING`,
		registration,
	)

	if err != nil {
		return err
	}

	return nil
}

// FindByNumber returns registrations with specified number.
func (r *RegistrationRepository) FindByNumber(number string) ([]model.Registration, error) {
	registrations := make([]model.Registration, 0)

	err := r.store.db.Select(&registrations,
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

	for i, reg := range registrations {
		if registrations[i].Date != nil {
			*registrations[i].Date = (*reg.Date)[:10]
		}

		if registrations[i].FirstRegDate != nil {
			*registrations[i].FirstRegDate = (*reg.FirstRegDate)[:10]
		}
	}

	return registrations, nil
}

// FindByCode returns registrations with specified code.
func (r *RegistrationRepository) FindByCode(code string) (*model.Registration, error) {
	var registration model.Registration

	err := r.store.db.Get(&registration,
		`SELECT brand, capacity, color, d_first_reg, d_reg, fuel, kind,
			    make_year, model, CONCAT(s_doc, n_doc) as code, n_reg_new, n_seating,
			    n_standing, own_weight, rank_category, total_weight, vin
		FROM registrations
		WHERE CONCAT(s_doc, n_doc) = $1
		ORDER BY d_reg DESC`,
		code,
	)

	if err == sql.ErrNoRows {
		return nil, store.ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	if registration.Date != nil {
		*registration.Date = (*registration.Date)[:10]
	}

	if registration.FirstRegDate != nil {
		*registration.FirstRegDate = (*registration.FirstRegDate)[:10]
	}

	return &registration, nil
}

// FindByVIN returns registrations with specified VIN.
func (r *RegistrationRepository) FindByVIN(vin string) ([]model.Registration, error) {
	registrations := make([]model.Registration, 0)

	err := r.store.db.Select(&registrations,
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

	for i, reg := range registrations {
		if registrations[i].Date != nil && len(*reg.FirstRegDate) >= 10 {
			*registrations[i].Date = (*reg.Date)[:10]
		}

		if registrations[i].FirstRegDate != nil && len(*reg.FirstRegDate) >= 10 {
			*registrations[i].FirstRegDate = (*reg.FirstRegDate)[:10]
		}
	}

	return registrations, nil
}

// GetLast returns last registration from the database.
func (r *RegistrationRepository) GetLast(series string) (*model.Registration, error) {
	var registration model.Registration

	err := r.store.db.Get(&registration,
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

	if err == sql.ErrNoRows {
		return nil, store.ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	if registration.Date != nil {
		*registration.Date = (*registration.Date)[:10]
	}

	if registration.FirstRegDate != nil {
		*registration.FirstRegDate = (*registration.FirstRegDate)[:10]
	}

	return &registration, nil
}

// AllSeries returns list of all known series.
func (r *RegistrationRepository) AllSeries() ([]string, error) {
	codes := make([]string, 0)

	err := r.store.db.Select(&codes, `SELECT s_doc FROM registrations GROUP BY s_doc`)
	if err == sql.ErrNoRows {
		return nil, store.ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	return codes, nil
}

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
		)`,
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
			    make_year, model, n_doc, s_doc, n_reg_new, n_seating,
			    n_standing, own_weight, rank_category, total_weight, vin
		FROM registrations
		WHERE n_reg_new = $1`,
		number,
	)

	if err != nil {
		return nil, err
	}

	for i, reg := range registrations {
		if registrations[i].DReg != nil {
			*registrations[i].DReg = (*reg.DReg)[:10]
		}

		if registrations[i].DFirstReg != nil {
			*registrations[i].DFirstReg = (*reg.DFirstReg)[:10]
		}
	}

	return registrations, nil
}

// FindByCode returns registrations with specified code.
func (r *RegistrationRepository) FindByCode(code string) (*model.Registration, error) {
	var registration model.Registration

	err := r.store.db.Get(&registration,
		`SELECT brand, capacity, color, d_first_reg, d_reg, fuel, kind,
			    make_year, model, CONCAT(n_doc, s_doc) as code, n_reg_new, n_seating,
			    n_standing, own_weight, rank_category, total_weight, vin
		FROM registrations
		WHERE code = $1`,
		code,
	)

	if err == sql.ErrNoRows {
		return nil, store.ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	if registration.DReg != nil {
		*registration.DReg = (*registration.DReg)[:10]
	}

	if registration.DFirstReg != nil {
		*registration.DFirstReg = (*registration.DFirstReg)[:10]
	}

	return &registration, nil
}

// FindByVIN returns registrations with specified VIN.
func (r *RegistrationRepository) FindByVIN(vin string) ([]model.Registration, error) {
	registrations := make([]model.Registration, 0)

	err := r.store.db.Select(&registrations,
		`SELECT brand, capacity, color, d_first_reg, d_reg, fuel, kind,
                make_year, model, CONCAT(n_doc, s_doc) as code, n_reg_new, n_seating,
                n_standing, own_weight, rank_category, total_weight, vin
		FROM registrations
		WHERE vin = $1`,
		vin,
	)

	for i, reg := range registrations {
		if registrations[i].DReg != nil {
			*registrations[i].DReg = (*reg.DReg)[:10]
		}

		if registrations[i].DFirstReg != nil {
			*registrations[i].DFirstReg = (*reg.DFirstReg)[:10]
		}
	}

	if err != nil {
		return nil, err
	}

	return registrations, nil
}

// GetLast returns last registration from the database.
func (r *RegistrationRepository) GetLast(series string) (*model.Registration, error) {
	var registration model.Registration

	err := r.store.db.Get(&registration,
		`SELECT
			brand, capacity, color, d_first_reg, d_reg, fuel,
			kind, make_year, model, CONCAT(n_doc, s_doc) as code, n_reg_new, n_seating,
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

	if registration.DReg != nil {
		*registration.DReg = (*registration.DReg)[:10]
	}

	if registration.DFirstReg != nil {
		*registration.DFirstReg = (*registration.DFirstReg)[:10]
	}

	return &registration, nil
}

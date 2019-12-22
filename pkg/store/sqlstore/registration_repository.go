package sqlstore

import (
	"database/sql"

	"github.com/opencars/edrmvs/pkg/model"
	"github.com/opencars/edrmvs/pkg/store"
)

type RegistrationRepository struct {
	store *Store
}

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

func (r *RegistrationRepository) FindByNumber(number string) ([]model.Registration, error) {
	registrations := make([]model.Registration, 0)

	err := r.store.db.Select(&registrations,
		`SELECT (
			brand, capacity, color, d_first_reg, d_reg, fuel, kind,
			make_year, model, n_doc, s_doc, n_reg_new, n_seating,
			n_standing, own_weight, rank_category, total_weight, vin
		)
		FROM registrations
		WHERE n_reg_new = $1`,
		number,
	)

	if err != nil {
		return nil, err
	}

	return registrations, nil
}

func (r *RegistrationRepository) FindByCode(code string) (*model.Registration, error) {
	var registration model.Registration
	err := r.store.db.Get(&registration,
		`SELECT (
			brand, capacity, color, d_first_reg, d_reg, fuel, kind,
			make_year, model, n_doc, s_doc, n_reg_new, n_seating,
			n_standing, own_weight, rank_category, total_weight, vin
		)
		FROM registrations
		WHERE CONCAT(s_doc, n_doc) = $1`,
		code,
	)

	if err == sql.ErrNoRows {
		return nil, store.RecordNotFound
	}

	if err != nil {
		return nil, err
	}

	return &registration, nil
}

func (r *RegistrationRepository) FindByVIN(vin string) ([]model.Registration, error) {
	registrations := make([]model.Registration, 0)

	err := r.store.db.Select(&registrations,
		`SELECT (
			brand, capacity, color, d_first_reg, d_reg, fuel, kind,
			make_year, model, n_doc, s_doc, n_reg_new, n_seating,
			n_standing, own_weight, rank_category, total_weight, vin
		)
		FROM registrations
		WHERE vin = $1`,
		vin,
	)

	if err != nil {
		return nil, err
	}

	return registrations, nil
}

func (r *RegistrationRepository) GetLast(series string) (*model.Registration, error) {
	var registration model.Registration

	err := r.store.db.Get(&registration,
		`SELECT
			brand, capacity, color, d_first_reg, d_reg, fuel,
			kind, make_year, model, n_doc, s_doc, n_reg_new, n_seating,
			n_standing, own_weight, rank_category, total_weight, vin
		FROM registrations
		WHERE s_doc = $1
		ORDER BY s_doc, n_doc DESC LIMIT 1`,
		series,
	)

	if err == sql.ErrNoRows {
		return nil, store.RecordNotFound
	}

	if err != nil {
		return nil, err
	}

	return &registration, nil
}

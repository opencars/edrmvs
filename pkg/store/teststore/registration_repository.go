package teststore

import (
	"github.com/opencars/edrmvs/pkg/model"
	"github.com/opencars/edrmvs/pkg/store"
)

// RegistrationRepository is responsible for registrations manipulation.
type RegistrationRepository struct {
	data []model.Registration
}

// Create adds new record to the database.
func (r *RegistrationRepository) Create(registration *model.Registration) error {
	r.data = append(r.data, *registration)
	return nil
}

// FindByNumber returns registrations with specified number.
func (r *RegistrationRepository) FindByNumber(number string) ([]model.Registration, error) {
	result := make([]model.Registration, 0)
	for i := range r.data {
		if r.data[i].Number == number {
			result = append(result, r.data[i])
		}
	}

	return result, nil
}

// FindByCode returns registrations with specified code.
func (r *RegistrationRepository) FindByCode(code string) (*model.Registration, error) {
	for _, registration := range r.data {
		if registration.SDoc+registration.NDoc == code {
			return &registration, nil
		}
	}

	return nil, store.ErrRecordNotFound
}

// FindByVIN returns registrations with specified VIN.
func (r *RegistrationRepository) FindByVIN(vin string) ([]model.Registration, error) {
	result := make([]model.Registration, 0)
	for i := range r.data {
		if r.data[i].VIN != nil && *r.data[i].VIN == vin {
			result = append(result, r.data[i])
		}
	}

	return result, nil
}

// GetLast returns last registration from the database.
func (r *RegistrationRepository) GetLast(series string) (*model.Registration, error) {
	for i := len(r.data) - 1; i >= 0; i-- {
		if r.data[i].SDoc == series {
			return &r.data[i], nil
		}
	}

	return nil, store.ErrRecordNotFound
}

// AllSeries returns list of all known series.
func (r *RegistrationRepository) AllSeries() ([]string, error) {
	series := make(map[string]struct{}, 0)

	for _, v := range r.data {
		if _, ok := series[v.SDoc]; !ok {
			series[v.SDoc] = struct{}{}
		}
	}

	result := make([]string, 0, len(series))
	for k := range series {
		result = append(result, k)
	}

	return result, nil
}

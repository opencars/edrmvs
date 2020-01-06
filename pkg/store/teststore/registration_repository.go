package teststore

import (
	"github.com/opencars/edrmvs/pkg/model"
)

// RegistrationRepository is responsible for registrations manipulation.
type RegistrationRepository struct {
	data map[string]*model.Registration
}

// Create adds new record to the database.
func (r *RegistrationRepository) Create(registration *model.Registration) error {
	r.data[registration.SDoc+registration.NDoc] = registration

	return nil
}

// FindByNumber returns registrations with specified number.
func (r *RegistrationRepository) FindByNumber(number string) ([]model.Registration, error) {
	return nil, nil
}

// FindByCode returns registrations with specified code.
func (r *RegistrationRepository) FindByCode(code string) (*model.Registration, error) {
	return nil, nil
}

// FindByVIN returns registrations with specified VIN.
func (r *RegistrationRepository) FindByVIN(vin string) ([]model.Registration, error) {
	return nil, nil
}

// GetLast returns last registration from the database.
func (r *RegistrationRepository) GetLast(series string) (*model.Registration, error) {
	return nil, nil
}

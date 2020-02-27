package store

import (
	"github.com/opencars/edrmvs/pkg/model"
)

// RegistrationRepository is responsible for registrations manipulation.
type RegistrationRepository interface {
	Create(registration *model.Registration) error
	FindByNumber(number string) ([]model.Registration, error)
	FindByCode(code string) (*model.Registration, error)
	FindByVIN(vin string) ([]model.Registration, error)
	GetLast(series string) (*model.Registration, error)
	AllSeries() ([]string, error)
}

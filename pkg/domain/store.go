package domain

import (
	"context"

	"github.com/opencars/edrmvs/pkg/domain/model"
)

type RegistrationStore interface {
	FindByCode(ctx context.Context, code string) (*model.Registration, error)
	FindByVIN(ctx context.Context, vin string) ([]model.Registration, error)
	FindByNumber(ctx context.Context, number string) ([]model.Registration, error)
	Health(ctx context.Context) error
}

type SystemRegistrationStore interface {
	Create(ctx context.Context, registration *model.Registration) error
	FindLastBySeries(ctx context.Context, series string) (*model.Registration, error)
	AllSeries(ctx context.Context) ([]string, error)
}

type FullRegistrationStore interface {
	SystemRegistrationStore
	RegistrationStore
}

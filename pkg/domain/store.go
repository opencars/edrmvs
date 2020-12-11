package domain

import (
	"context"
)

type RegistrationStore interface {
	FindByCode(ctx context.Context, code string) (*Registration, error)
	FindByVIN(ctx context.Context, vin string) ([]Registration, error)
	FindByNumber(ctx context.Context, number string) ([]Registration, error)
	Health(ctx context.Context) error
}

type SystemRegistrationStore interface {
	Create(ctx context.Context, registration *Registration) error
	FindLastBySeries(ctx context.Context, series string) (*Registration, error)
	AllSeries(ctx context.Context) ([]string, error)
}

type FullRegistrationStore interface {
	SystemRegistrationStore
	RegistrationStore
}

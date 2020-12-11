package domain

import (
	"context"
)

type RegistrationService interface {
	FindByNumber(ctx context.Context, number string) ([]Registration, error)
	FindByCode(ctx context.Context, code string) (*Registration, error)
	FindByVIN(ctx context.Context, vin string) ([]Registration, error)
	Health(ctx context.Context) error
}

type RegistrationProvider interface {
	FindByCode(ctx context.Context, code string) ([]Registration, error)
}

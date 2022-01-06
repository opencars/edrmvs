package domain

import (
	"context"

	"github.com/opencars/edrmvs/pkg/domain/model"
)

type RegistrationService interface {
	FindByNumber(ctx context.Context, number string) ([]model.Registration, error)
	FindByCode(ctx context.Context, code string) (*model.Registration, error)
	FindByVIN(ctx context.Context, vin string, v2 bool) ([]model.Registration, error)
	Health(ctx context.Context) error
}

type RegistrationProvider interface {
	FindByCode(ctx context.Context, code string) ([]model.Registration, error)
}

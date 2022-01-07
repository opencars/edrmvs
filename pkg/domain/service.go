package domain

import (
	"context"

	"github.com/opencars/edrmvs/pkg/domain/model"
	"github.com/opencars/edrmvs/pkg/domain/query"
)

type CustomerService interface {
	ListByNumber(ctx context.Context, q *query.ListByNumber) ([]model.Registration, error)
	ListByVIN(ctx context.Context, q *query.ListByVIN, v2 bool) ([]model.Registration, error)
	DetailsByCode(ctx context.Context, q *query.DetailsByCode) (*model.Registration, error)
	Health(ctx context.Context) error
}

type InternalService interface {
	ListByNumber(ctx context.Context, q *query.ListWithNumberByInternal) ([]model.Registration, error)
	ListByVIN(ctx context.Context, q *query.ListWithVINByInternal) ([]model.Registration, error)
	DetailsByCode(ctx context.Context, q *query.DetailsWithCodeByInternal) (*model.Registration, error)
}

type RegistrationProvider interface {
	FindByCode(ctx context.Context, code string) ([]model.Registration, error)
}

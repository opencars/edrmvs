package service

import (
	"context"

	"github.com/opencars/translit"

	"github.com/opencars/edrmvs/pkg/domain"
	"github.com/opencars/edrmvs/pkg/domain/model"
	"github.com/opencars/edrmvs/pkg/domain/query"
	"github.com/opencars/edrmvs/pkg/logger"
)

type InternalService struct {
	s domain.RegistrationStore
	p domain.RegistrationProvider
}

func NewInternalService(s domain.RegistrationStore, p domain.RegistrationProvider) *InternalService {
	return &InternalService{
		s: s,
		p: p,
	}
}

func (s *InternalService) ListByNumber(ctx context.Context, q *query.ListWithNumberByInternal) ([]model.Registration, error) {
	if err := query.Process(q); err != nil {
		return nil, err
	}

	return s.s.FindByNumber(ctx, q.Number)
}

func (s *InternalService) ListByVIN(ctx context.Context, q *query.ListWithVINByInternal) ([]model.Registration, error) {
	if err := query.Process(q); err != nil {
		return nil, err
	}

	registrations, err := s.s.FindByVIN(ctx, q.VIN)
	if err != nil {
		return nil, err
	}

	for i := range registrations {
		registrations[i].Code = translit.ToLatin(registrations[i].Code)

		var isActive bool
		items, err := s.p.FindByCode(ctx, registrations[i].Code)
		if err != nil {
			logger.Errorf("failed to check is_active: %s", err)
			continue
		}

		for _, item := range items {
			if item.Number == registrations[i].Number {
				isActive = true
			}
		}

		registrations[i].IsActive = &isActive
	}

	return registrations, nil
}

func (s *InternalService) DetailsByCode(ctx context.Context, q *query.DetailsWithCodeByInternal) (*model.Registration, error) {
	if err := query.Process(q); err != nil {
		return nil, err
	}

	return s.s.FindByCode(ctx, q.Code)
}

package service

import (
	"context"
	"strings"

	"github.com/opencars/schema"
	"github.com/opencars/translit"

	"github.com/opencars/edrmvs/pkg/domain"
	"github.com/opencars/edrmvs/pkg/domain/model"
	"github.com/opencars/edrmvs/pkg/domain/query"
	"github.com/opencars/edrmvs/pkg/logger"
)

type CustomerService struct {
	s        domain.RegistrationStore
	p        domain.RegistrationProvider
	producer schema.Producer
}

func NewCustomerService(s domain.RegistrationStore, p domain.RegistrationProvider, producer schema.Producer) *CustomerService {
	return &CustomerService{
		s:        s,
		p:        p,
		producer: producer,
	}
}

func (s *CustomerService) ListByNumber(ctx context.Context, q *query.ListByNumber) ([]model.Registration, error) {
	number := translit.ToLatin(strings.ToUpper(q.Number))

	registrations, err := s.s.FindByNumber(ctx, number)
	if err != nil {
		return nil, err
	}

	if err := s.producer.Produce(ctx, q.Event(registrations...)); err != nil {
		return nil, err
	}

	return registrations, nil
}

func (s *CustomerService) ListByVIN(ctx context.Context, q *query.ListByVIN, v2 bool) ([]model.Registration, error) {
	vin := translit.ToLatin(strings.ToUpper(q.VIN))

	registrations, err := s.s.FindByVIN(ctx, vin)
	if err != nil {
		return nil, err
	}

	if !v2 {
		if err := s.producer.Produce(ctx, q.Event(registrations...)); err != nil {
			return nil, err
		}

		return registrations, nil
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

	if err := s.producer.Produce(ctx, q.Event(registrations...)); err != nil {
		return nil, err
	}

	return registrations, nil
}

func (s *CustomerService) DetailsByCode(ctx context.Context, q *query.DetailsByCode) (*model.Registration, error) {
	code := translit.ToLatin(strings.ToUpper(q.Code))

	registration, err := s.s.FindByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	if err := s.producer.Produce(ctx, q.Event(registration)); err != nil {
		return nil, err
	}

	return registration, nil
}

// TODO: Move out of the user service:
func (s *CustomerService) Health(ctx context.Context) error {
	return s.s.Health(ctx)
}
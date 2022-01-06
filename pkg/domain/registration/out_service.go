package registration

import (
	"context"
	"strings"

	"github.com/opencars/translit"

	"github.com/opencars/edrmvs/pkg/domain"
	"github.com/opencars/edrmvs/pkg/domain/model"
	"github.com/opencars/edrmvs/pkg/logger"
)

type OutService struct {
	s domain.RegistrationStore
	p domain.RegistrationProvider
}

func NewService(s domain.RegistrationStore, p domain.RegistrationProvider) *OutService {
	return &OutService{
		s: s,
		p: p,
	}
}

func (svc *OutService) FindByNumber(ctx context.Context, lexeme string) ([]model.Registration, error) {
	number := translit.ToLatin(strings.ToUpper(lexeme))

	if len(number) < 6 {
		return nil, domain.ErrBadNumber
	}

	return svc.s.FindByNumber(ctx, number)
}

func (svc *OutService) FindByVIN(ctx context.Context, lexeme string, v2 bool) ([]model.Registration, error) {
	vin := translit.ToLatin(strings.ToUpper(lexeme))

	if len(vin) < 6 {
		return nil, domain.ErrBadVIN
	}

	registrations, err := svc.s.FindByVIN(ctx, vin)
	if err != nil {
		return nil, err
	}

	if !v2 {
		return registrations, nil
	}

	for i := range registrations {
		registrations[i].Code = translit.ToLatin(registrations[i].Code)

		var isActive bool
		items, err := svc.p.FindByCode(ctx, registrations[i].Code)
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

func (svc *OutService) FindByCode(ctx context.Context, lexeme string) (*model.Registration, error) {
	code := translit.ToLatin(strings.ToUpper(lexeme))
	if len(code) != 9 {
		return nil, domain.ErrBadCode
	}

	return svc.s.FindByCode(ctx, code)
}

func (svc *OutService) Health(ctx context.Context) error {
	return svc.s.Health(ctx)
}

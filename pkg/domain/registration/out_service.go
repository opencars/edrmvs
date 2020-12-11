package registration

import (
	"context"
	"strings"

	"github.com/opencars/edrmvs/pkg/domain"
	"github.com/opencars/translit"
)

type OutService struct {
	store domain.RegistrationStore
}

func NewRegistrationService(store domain.RegistrationStore) *OutService {
	return &OutService{
		store: store,
	}
}

func (svc *OutService) FindByNumber(ctx context.Context, lexeme string) ([]domain.Registration, error) {
	number := translit.ToLatin(strings.ToUpper(lexeme))

	if len(number) < 6 {
		return nil, domain.ErrBadNumber
	}

	return svc.store.FindByNumber(ctx, number)
}

func (svc *OutService) FindByVIN(ctx context.Context, lexeme string) ([]domain.Registration, error) {
	vin := translit.ToLatin(strings.ToUpper(lexeme))

	if len(vin) < 6 {
		return nil, domain.ErrBadVIN
	}

	return svc.store.FindByVIN(ctx, vin)
}

func (svc *OutService) FindByCode(ctx context.Context, lexeme string) (*domain.Registration, error) {
	code := translit.ToLatin(strings.ToUpper(lexeme))
	if len(code) != 9 {
		return nil, domain.ErrBadCode
	}

	return svc.store.FindByCode(ctx, code)
}

func (svc *OutService) Health(ctx context.Context) error {
	return svc.store.Health(ctx)
}

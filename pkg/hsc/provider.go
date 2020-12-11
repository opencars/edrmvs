package hsc

import (
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/opencars/edrmvs/pkg/domain"
)

var (
	Petrol      = "БЕНЗИН"
	Gas         = "ГАЗ"
	PetrolOrGas = "БЕНЗИН АБО ГАЗ"
	Diesel      = "ДИЗЕЛЬНЕ ПАЛИВО"
)

type Provider struct {
	api *API

	m       sync.RWMutex
	session *Session
	expAt   time.Time
}

func NewProvider(api *API) *Provider {
	return &Provider{
		api: api,
	}
}

func (p *Provider) FindByCode(ctx context.Context, code string) ([]domain.Registration, error) {
	p.m.Lock()
	if !p.expAt.After(time.Now()) {

		session, err := p.api.Authorize(ctx)
		if err != nil {
			return nil, err
		}

		p.expAt = time.Now().Add(time.Second * time.Duration(session.ExpiresIn))
		p.session = session
	}
	p.m.Unlock()

	p.m.RLock()
	defer p.m.RUnlock()

	registrations, err := p.api.VehiclePassport(ctx, p.session.AccessToken, code)
	if err != nil {
		return nil, err
	}

	result := make([]domain.Registration, 0, len(registrations))

	for i := range registrations {
		dto, err := convert(&registrations[i])
		if err != nil {
			return nil, err
		}

		result = append(result, *dto)
	}

	return result, nil
}

func convert(r *Registration) (*domain.Registration, error) {
	dto := domain.Registration{
		Brand:          r.Brand,
		Color:          r.Color,
		FirstRegDate:   r.FirstDate,
		Date:           r.Date,
		Kind:           &r.CommercialDesc,
		Year:           r.MakeYear,
		Model:          r.Model,
		DocumentSeries: r.Series,
		DocumentNumber: r.Number,
		Code:           r.Series + r.Number,
		Number:         r.LicencePlate,
		RankCategory:   r.Category,
		VIN:            r.NChassis,
	}

	dto.Fuel = fuel(r.Fuel)

	if r.Capacity != nil {
		capacity, err := strconv.Atoi(*r.Capacity)
		if err != nil {
			return nil, err
		}

		dto.Capacity = &capacity
	}

	if r.NSeating != nil {
		numSeating, err := strconv.Atoi(*r.NSeating)
		if err != nil {
			return nil, err
		}

		dto.NumSeating = &numSeating
	}

	if r.NStanding != nil {
		numStanding, err := strconv.Atoi(*r.NStanding)
		if err != nil {
			return nil, err
		}

		dto.NumStanding = &numStanding
	}

	if r.OwnWeight != nil {
		ownWeight, err := strconv.ParseFloat(*r.OwnWeight, 64)
		if err != nil {
			return nil, err
		}

		dto.OwnWeight = &ownWeight
	}

	if r.TotalWeight != nil {
		totalWeight, err := strconv.ParseFloat(*r.TotalWeight, 64)
		if err != nil {
			return nil, err
		}

		dto.TotalWeight = &totalWeight
	}

	return &dto, nil
}

func fuel(fuel *string) *string {
	if fuel != nil {
		switch *fuel {
		case "B":
			return &Petrol
		case "S":
			return &PetrolOrGas
		case "G":
			return &Gas
		case "D":
			return &Diesel
		}
	}

	return nil
}

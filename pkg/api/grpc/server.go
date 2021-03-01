package grpc

import (
	"context"

	"github.com/opencars/grpc/pkg/registration"
)

type registrationHandler struct {
	registration.UnimplementedServiceServer
	api *API
}

func (h *registrationHandler) FindByNumber(ctx context.Context, r *registration.NumberRequest) (*registration.Response, error) {
	registrations, err := h.api.svc.FindByNumber(ctx, r.Number)
	if err != nil {
		return nil, handleErr(err)
	}

	dto := registration.Response{
		Registrations: make([]*registration.Record, 0, len(registrations)),
	}

	for i := range registrations {
		dto.Registrations = append(dto.Registrations, FromDomain(&registrations[i]))
	}

	return &dto, nil
}

func (h *registrationHandler) FindByVIN(ctx context.Context, r *registration.VINRequest) (*registration.Response, error) {
	registrations, err := h.api.svc.FindByVIN(ctx, r.Vin, true)
	if err != nil {
		return nil, handleErr(err)
	}

	dto := registration.Response{
		Registrations: make([]*registration.Record, 0, len(registrations)),
	}

	for i := range registrations {
		dto.Registrations = append(dto.Registrations, FromDomain(&registrations[i]))
	}

	return &dto, nil
}

func (h *registrationHandler) FindByCode(ctx context.Context, r *registration.CodeRequest) (*registration.Response, error) {
	object, err := h.api.svc.FindByCode(ctx, r.Code)
	if err != nil {
		return nil, handleErr(err)
	}

	return &registration.Response{
		Registrations: []*registration.Record{FromDomain(object)},
	}, nil
}

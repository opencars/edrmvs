package grpc

import (
	"context"

	"github.com/opencars/grpc/pkg/registration"

	"github.com/opencars/edrmvs/pkg/domain/query"
)

type registrationHandler struct {
	registration.UnimplementedServiceServer
	api *API
}

func (h *registrationHandler) FindByNumber(ctx context.Context, r *registration.NumberRequest) (*registration.Response, error) {
	q := query.ListWithNumberByInternal{
		Number: r.GetNumber(),
	}

	registrations, err := h.api.svc.ListByNumber(ctx, &q)
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
	q := query.ListWithVINByInternal{
		VIN: r.GetVin(),
	}

	registrations, err := h.api.svc.ListByVIN(ctx, &q)
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
	q := query.DetailsWithCodeByInternal{
		Code: r.GetCode(),
	}

	object, err := h.api.svc.DetailsByCode(ctx, &q)
	if err != nil {
		return nil, handleErr(err)
	}

	return &registration.Response{
		Registrations: []*registration.Record{
			FromDomain(object),
		},
	}, nil
}

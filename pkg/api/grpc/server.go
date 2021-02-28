package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/opencars/grpc/pkg/registration"
)

type registrationHandler struct {
	registration.UnimplementedVehicleServiceServer
	api *API
}

func (h *registrationHandler) FindByNumber(ctx context.Context, r *registration.FindByNumberRequest) (*registration.RegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByNumber not implemented")
}

func (h *registrationHandler) FindByVIN(ctx context.Context, r *registration.FindByVINRequest) (*registration.RegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByVIN not implemented")
}

func (h *registrationHandler) FindByCode(ctx context.Context, r *registration.FindByCodeRequest) (*registration.RegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByCode not implemented")
}

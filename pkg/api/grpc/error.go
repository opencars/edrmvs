package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/opencars/edrmvs/pkg/domain"
)

var (
	ErrNotFound  = status.Error(codes.NotFound, "record.not_found")
	ErrBadCode   = status.Error(codes.InvalidArgument, "request.bad_code")
	ErrBadVIN    = status.Error(codes.InvalidArgument, "request.bad_vin")
	ErrBadNumber = status.Error(codes.InvalidArgument, "request.bad_number")
)

func handleErr(err error) error {
	switch err {
	case domain.ErrNotFound:
		return ErrNotFound
	case domain.ErrBadCode:
		return domain.ErrBadCode
	case domain.ErrBadVIN:
		return domain.ErrBadVIN
	case domain.ErrBadNumber:
		return domain.ErrBadNumber
	default:
		return err
	}
}

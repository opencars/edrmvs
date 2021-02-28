package grpc

import (
	"net"

	"github.com/opencars/grpc/pkg/registration"
	"google.golang.org/grpc"

	"github.com/opencars/edrmvs/pkg/domain"
)

// API represents gRPC API server.
type API struct {
	Addr string
	s    *grpc.Server
	svc  domain.RegistrationService
}

func New(addr string, svc domain.RegistrationService, /* conf *config.GRPCServer */) *API {
	return &API{
		Addr: addr,
		svc:  svc,
	}
}

func (a *API) Run() error {
	listener, err := net.Listen("tcp", a.Addr)
	if err != nil {
		return err
	}

	// TODO: Add logging & metrics & errors.
	//opts := []grpc.ServerOption{
	//	grpc.ChainUnaryInterceptor(
	//		RequestLoggingInterceptor,
	//		ErrorInterceptor,
	//	),
	//}

	a.s = grpc.NewServer()
	registration.RegisterVehicleServiceServer(a.s, &registrationHandler{api: a})

	return a.s.Serve(listener)
}

// Close gracefully stops grpc API server.
func (a *API) Close() error {
	if a.s != nil {
		a.s.GracefulStop()
	}

	return nil
}

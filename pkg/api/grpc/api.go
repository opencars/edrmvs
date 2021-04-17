package grpc

import (
	"context"
	"net"

	"github.com/opencars/grpc/pkg/registration"
	"google.golang.org/grpc"

	"github.com/opencars/edrmvs/pkg/domain"
)

// API represents gRPC API server.
type API struct {
	addr string
	s    *grpc.Server
	svc  domain.RegistrationService
}

func New(addr string, svc domain.RegistrationService) *API {
	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			RequestLoggingInterceptor,
		),
	}

	return &API{
		addr: addr,
		svc:  svc,
		s:    grpc.NewServer(opts...),
	}
}

func (a *API) Run(ctx context.Context) error {
	listener, err := net.Listen("tcp", a.addr)
	if err != nil {
		return err
	}
	defer listener.Close()

	registration.RegisterServiceServer(a.s, &registrationHandler{api: a})

	errors := make(chan error)
	go func() {
		errors <- a.s.Serve(listener)
	}()

	select {
	case <-ctx.Done():
		a.s.GracefulStop()
		return <-errors
	case err := <-errors:
		return err
	}
}

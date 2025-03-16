package grpc

import (
	"context"
	"net"

	"github.com/opencars/grpc/pkg/registration"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/opencars/edrmvs/pkg/domain"
)

// API represents gRPC API server.
type API struct {
	addr string
	s    *grpc.Server
	svc  domain.InternalService
}

func New(addr string, svc domain.InternalService) *API {
	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			RequestLoggingInterceptor,
		),
	}

	server := grpc.NewServer(opts...)
	// Enable reflection for debugging and development tools
	reflection.Register(server)

	return &API{
		addr: addr,
		svc:  svc,
		s:    server,
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

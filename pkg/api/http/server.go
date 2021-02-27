package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"runtime"

	"github.com/gorilla/mux"

	"github.com/opencars/edrmvs/pkg/domain"
	"github.com/opencars/edrmvs/pkg/handler"
	"github.com/opencars/edrmvs/pkg/version"
)

type server struct {
	router *mux.Router

	svc domain.RegistrationService
}

func newServer(svc domain.RegistrationStore) *server {
	srv := server{
		router: mux.NewRouter(),
		svc:    svc,
	}

	srv.configureRouter()

	return &srv
}

func (*server) Version() handler.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		v := struct {
			Version string `json:"version"`
			Go      string `json:"go"`
		}{
			Version: version.Version,
			Go:      runtime.Version(),
		}

		return json.NewEncoder(w).Encode(v)
	}
}

func (s *server) Health() handler.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		if err := s.svc.Health(r.Context()); err != nil {
			return err
		}

		return nil
	}
}

func (s *server) FindByVIN() handler.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		registrations, err := s.svc.FindByVIN(r.Context(), r.URL.Query().Get("vin"))
		if errors.Is(err, domain.ErrBadVIN) {
			return handler.ErrBadVIN
		}

		if err != nil {
			return err
		}

		return json.NewEncoder(w).Encode(registrations)
	}
}

func (s *server) FindByNumber() handler.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		registrations, err := s.svc.FindByNumber(r.Context(), r.URL.Query().Get("number"))
		if errors.Is(err, domain.ErrBadNumber) {
			return handler.ErrBadNumber
		}

		if err != nil {
			return err
		}

		return json.NewEncoder(w).Encode(registrations)
	}
}

func (s *server) FindByCode() handler.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		registration, err := s.svc.FindByCode(r.Context(), mux.Vars(r)["code"])
		if errors.Is(err, domain.ErrBadCode) {
			return handler.ErrBadCode
		}

		if err != nil {
			return err
		}

		return json.NewEncoder(w).Encode(registration)
	}
}

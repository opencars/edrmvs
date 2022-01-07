package http

import (
	"encoding/json"
	"net/http"
	"runtime"

	"github.com/gorilla/mux"
	"github.com/opencars/httputil"

	"github.com/opencars/edrmvs/pkg/domain"
	"github.com/opencars/edrmvs/pkg/domain/query"
	"github.com/opencars/edrmvs/pkg/version"
)

type server struct {
	router *mux.Router

	svc domain.CustomerService
}

func newServer(svc domain.CustomerService) *server {
	srv := server{
		router: mux.NewRouter(),
		svc:    svc,
	}

	srv.configureRouter()

	return &srv
}

func (*server) Version() httputil.Handler {
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

func (s *server) Health() httputil.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		if err := s.svc.Health(r.Context()); err != nil {
			return err
		}

		return nil
	}
}

func (s *server) FindByVIN(v2 bool) httputil.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		q := query.ListByVIN{
			UserID: UserIDFromContext(r.Context()),
			VIN:    r.URL.Query().Get("vin"),
		}

		registrations, err := s.svc.ListByVIN(r.Context(), &q, v2)
		if err != nil {
			return handleErr(err)
		}

		return json.NewEncoder(w).Encode(registrations)
	}
}

func (s *server) FindByNumber() httputil.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		q := query.ListByNumber{
			UserID: UserIDFromContext(r.Context()),
			Number: r.URL.Query().Get("number"),
		}

		registrations, err := s.svc.ListByNumber(r.Context(), &q)
		if err != nil {
			return handleErr(err)
		}

		return json.NewEncoder(w).Encode(registrations)
	}
}

func (s *server) FindByCode() httputil.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		q := query.DetailsByCode{
			UserID: UserIDFromContext(r.Context()),
			Code:   mux.Vars(r)["code"],
		}

		registration, err := s.svc.DetailsByCode(r.Context(), &q)
		if err != nil {
			return handleErr(err)
		}

		return json.NewEncoder(w).Encode(registration)
	}
}

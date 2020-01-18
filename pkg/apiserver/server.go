package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/opencars/edrmvs/pkg/store"
	"github.com/opencars/edrmvs/pkg/version"
)

func newServer(store store.Store) *server {
	srv := server{
		router: mux.NewRouter(),
		store:  store,
	}

	srv.confgureRoutes()

	return &srv
}

type server struct {
	router *mux.Router
	store  store.Store
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: Add CORS.

	s.router.ServeHTTP(w, r)
}

func (s *server) confgureRoutes() {
	s.router.Handle("/api/v1/edrmvs/version", version.Handler{})

	s.router.Handle("/api/v1/edrmvs/vin/{vin}", s.registrationByVIN())
	s.router.Handle("/api/v1/edrmvs/number/{number}", s.registrationByNumber())
	s.router.Handle("/api/v1/edrmvs/code/{code}", s.registrationByCode())
}

func (s *server) registrationByVIN() Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		vin := mux.Vars(r)["vin"]

		registrations, err := s.store.Registration().FindByVIN(vin)
		if err != nil {
			return err
		}

		if err := json.NewEncoder(w).Encode(registrations); err != nil {
			return err
		}

		return nil
	}
}

func (s *server) registrationByNumber() Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		number := mux.Vars(r)["number"]

		registrations, err := s.store.Registration().FindByNumber(number)
		if err != nil {
			return err
		}

		if err := json.NewEncoder(w).Encode(registrations); err != nil {
			return err
		}

		return nil
	}
}

func (s *server) registrationByCode() Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		code := mux.Vars(r)["code"]

		registration, err := s.store.Registration().FindByCode(code)
		if err != nil {
			return err
		}

		if err := json.NewEncoder(w).Encode(registration); err != nil {
			return err
		}

		return nil
	}
}

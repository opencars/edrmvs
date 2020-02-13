package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/opencars/edrmvs/pkg/store"
	"github.com/opencars/edrmvs/pkg/version"
	"github.com/opencars/translit"
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

// ServeHTTP implements http.Handler interface.
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"Api-Key"})

	cors := handlers.CORS(origins, methods, headers)(s.router)
	cors.ServeHTTP(w, r)
}

func (s *server) confgureRoutes() {
	s.router.Handle("/api/v1/registrations/version", version.Handler{}).Methods("GET", "OPTIONS")

	s.router.Handle("/api/v1/registrations", s.registrationByVIN()).Queries("vin", "{vin}").Methods("GET", "OPTIONS")
	s.router.Handle("/api/v1/registrations", s.registrationByNumber()).Queries("number", "{number}").Methods("GET", "OPTIONS")
	s.router.Handle("/api/v1/registrations/{code}", s.registrationByCode()).Methods("GET", "OPTIONS")
}

func (s *server) registrationByVIN() Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		vin := r.URL.Query().Get("vin")

		registrations, err := s.store.Registration().FindByVIN(vin)
		if err != nil {
			return err
		}

		for i, reg := range registrations {
			registrations[i].Code = reg.NDoc + reg.SDoc
		}

		if err := json.NewEncoder(w).Encode(registrations); err != nil {
			return err
		}

		return nil
	}
}

func (s *server) registrationByNumber() Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		number := translit.ToLatin(r.URL.Query().Get("number"))

		registrations, err := s.store.Registration().FindByNumber(number)
		if err != nil {
			return err
		}

		for i, reg := range registrations {
			registrations[i].Code = reg.NDoc + reg.SDoc
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

		registration.Code = registration.NDoc + registration.SDoc
		if err := json.NewEncoder(w).Encode(registration); err != nil {
			return err
		}

		return nil
	}
}

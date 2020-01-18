package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/handlers"
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

// ServeHTTP implements http.Handler interface.
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"Api-Key"})

	cors := handlers.CORS(origins, methods, headers)(s.router)
	cors.ServeHTTP(w, r)
}

func (s *server) confgureRoutes() {
	s.router.Handle("/api/v1/edrmvs/version", version.Handler{}).Methods("GET", "OPTIONS")

	s.router.Handle("/api/v1/edrmvs/vin/{vin}", s.registrationByVIN()).Methods("GET", "OPTIONS")
	s.router.Handle("/api/v1/edrmvs/number/{number}", s.registrationByNumber()).Methods("GET", "OPTIONS")
	s.router.Handle("/api/v1/edrmvs/code/{code}", s.registrationByCode()).Methods("GET", "OPTIONS")
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

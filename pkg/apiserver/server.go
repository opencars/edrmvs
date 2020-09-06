package apiserver

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/opencars/translit"

	"github.com/opencars/edrmvs/pkg/handler"
	"github.com/opencars/edrmvs/pkg/store"
)

func newServer(store store.Store) *server {
	srv := server{
		router: mux.NewRouter(),
		store:  store,
	}

	srv.configureRoutes()

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
	headers := handlers.AllowedHeaders([]string{"Api-Key", "X-Api-Key"})

	cors := handlers.CORS(origins, methods, headers)(s.router)
	compressed := handlers.CompressHandler(cors)
	compressed.ServeHTTP(w, r)
}

func (s *server) FindByVIN() handler.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		vin := r.URL.Query().Get("vin")
		if vin == "" {
			return handler.ErrNotFound
		}

		registrations, err := s.store.Registration().FindByVIN(vin)
		if err != nil {
			return err
		}

		return json.NewEncoder(w).Encode(registrations)
	}
}

func (s *server) FindByNumber() handler.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		number := translit.ToLatin(strings.ToUpper(r.URL.Query().Get("number")))

		registrations, err := s.store.Registration().FindByNumber(number)
		if err != nil {
			return err
		}

		return json.NewEncoder(w).Encode(registrations)
	}
}

func (s *server) FindByCode() handler.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		code := translit.ToLatin(strings.ToUpper(mux.Vars(r)["code"]))

		registration, err := s.store.Registration().FindByCode(code)
		if err == store.ErrRecordNotFound {
			return handler.ErrNotFound
		}

		if err != nil {
			return err
		}

		return json.NewEncoder(w).Encode(registration)
	}
}

package http

import (
	"encoding/json"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/opencars/translit"

	"github.com/opencars/edrmvs/pkg/handler"
	"github.com/opencars/edrmvs/pkg/store"
	"github.com/opencars/edrmvs/pkg/version"
)

const (
	// MaxImageSize equals to 5 MB.
	MaxImageSize = 5 << 20

	// ClientTimeOut equals to 10 seconds.
	ClientTimeOut = 10 * time.Second
)

type server struct {
	router *mux.Router
	store  store.Store
}

func newServer(store store.Store) *server {
	srv := server{
		router: mux.NewRouter(),
		store:  store,
	}

	srv.configureRouter()

	return &srv
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
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
		if err := s.store.Health(r.Context()); err != nil {
			return err
		}

		return nil
	}
}

func (s *server) FindByVIN() handler.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		vin := r.URL.Query().Get("vin")
		if len(vin) < 6 {
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
		if len(number) < 6 {
			return handler.ErrNotFound
		}

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
		if len(code) != 9 {
			return handler.ErrNotFound
		}

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

package apiserver

import (
	"net/http"

	"github.com/opencars/edrmvs/pkg/config"
	"github.com/opencars/edrmvs/pkg/store/sqlstore"
)

func Start(settings *config.Settings) error {
	store, err := sqlstore.New(settings)
	if err != nil {
		return err
	}

	server := newServer(store)
	srv := http.Server{
		Addr:    ":8080",
		Handler: server,
	}

	return srv.ListenAndServe()
}

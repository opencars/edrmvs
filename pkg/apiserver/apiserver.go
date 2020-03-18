package apiserver

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"

	"github.com/opencars/edrmvs/pkg/config"
	"github.com/opencars/edrmvs/pkg/store/sqlstore"
)

func Start(settings *config.Settings) error {
	store, err := sqlstore.New(&settings.DB)
	if err != nil {
		return err
	}

	server := newServer(store)
	srv := http.Server{
		Addr:    ":8080",
		Handler: handlers.LoggingHandler(os.Stdout, server),
	}

	log.Println("Server is listening...")
	return srv.ListenAndServe()
}

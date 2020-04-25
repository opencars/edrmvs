package apiserver

import (
	"github.com/opencars/edrmvs/pkg/version"
)

func (s *server) configureRoutes() {
	s.router.Handle("/api/v1/registrations/version", version.Handler{}).Methods("GET", "OPTIONS")

	s.router.Handle("/api/v1/registrations", s.FindByVIN()).Queries("vin", "{vin}").Methods("GET", "OPTIONS")
	s.router.Handle("/api/v1/registrations", s.FindByNumber()).Queries("number", "{number}").Methods("GET", "OPTIONS")
	s.router.Handle("/api/v1/registrations/{code}", s.FindByCode()).Methods("GET", "OPTIONS")
}

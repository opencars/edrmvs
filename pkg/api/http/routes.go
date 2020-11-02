package http

func (s *server) configureRouter() {
	router := s.router.PathPrefix("/api/v1/").Subrouter()

	router.Handle("/version", s.Version())
	router.Handle("/health", s.Health())

	router.Handle("/registrations", s.FindByVIN()).Queries("vin", "{vin}")
	router.Handle("/registrations", s.FindByNumber()).Queries("number", "{number}")
	router.Handle("/registrations/{code}", s.FindByCode())
}

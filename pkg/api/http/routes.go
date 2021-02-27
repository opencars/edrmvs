package http

func (s *server) configureRouter() {
	v1 := s.router.PathPrefix("/api/v1/").Subrouter()

	v1.Handle("/version", s.Version())
	v1.Handle("/health", s.Health())

	v1.Handle("/registrations", s.FindByVIN(false)).Queries("vin", "{vin}")
	v1.Handle("/registrations", s.FindByNumber()).Queries("number", "{number}")
	v1.Handle("/registrations/{code}", s.FindByCode())

	v2 := s.router.PathPrefix("/api/v2/").Subrouter()
	v2.Handle("/registrations", s.FindByVIN(true)).Queries("vin", "{vin}")
}

package http

func (s *server) configureRouter() {
	v1 := s.router.PathPrefix("/api/v1/").Subrouter()
	v2 := s.router.PathPrefix("/api/v2/").Subrouter()

	v1.Handle("/version", s.Version())
	v1.Handle("/health", s.Health())

	userv1 := v1.PathPrefix("").Subrouter()
	userv1.Use(
		AuthorizationMiddleware(),
	)

	userv1.Handle("/registrations", s.FindByVIN(false)).Queries("vin", "{vin}").Methods("GET")
	userv1.Handle("/registrations", s.FindByNumber()).Queries("number", "{number}").Methods("GET")
	userv1.Handle("/registrations/{code}", s.FindByCode()).Methods("GET")

	userv2 := v2.PathPrefix("").Subrouter()
	userv2.Use(
		AuthorizationMiddleware(),
	)

	userv2.Handle("/registrations", s.FindByVIN(true)).Queries("vin", "{vin}").Methods("GET")
}

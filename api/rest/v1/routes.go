package rest

import (
	tracingRouter "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
)

func (s *Server) routes() {
	s.Router = tracingRouter.NewRouter(tracingRouter.WithServiceName(s.ServiceName))

	s.Router.HandleFunc("/healthcheck", s.handleHealthcheck())
}

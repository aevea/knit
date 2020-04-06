package rest

import (
	"net/http"

	tracingRouter "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
)

// Server houses all dependencies and routing of the server
type Server struct {
	Router *tracingRouter.Router
	// ServiceName is used for tracing purposes
	ServiceName string
}

// NewServer creates a new instance of server and sets up routes
func NewServer(serviceName string) *Server {
	s := Server{
		ServiceName: serviceName,
	}
	s.routes()

	return &s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

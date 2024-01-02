// File: server.go
package pkg

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/shashankmahajan99/awesome-expense-tracker-backend/api"
)

// Server is the gRPC server.
type Server struct {
	api.UnimplementedUserAuthenticationServer
	router *runtime.ServeMux
}

// NewServer creates a new server.
func NewServer() (server *Server, err error) {
	err = server.Setup()
	if err != nil {
		return nil, err
	}
	return server, nil
}

// Setup routes.
func (s *Server) Setup() error {
	router := runtime.NewServeMux()
	s.router = router
	return nil
}

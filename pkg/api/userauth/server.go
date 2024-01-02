// File: server.go
package apipkg

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	AwesomeExpenseTrackerApi "github.com/shashankmahajan99/awesome-expense-tracker-backend/api"
)

// Server is the gRPC server.
type Server struct {
	AwesomeExpenseTrackerApi.UnimplementedUserAuthenticationServer
	router *runtime.ServeMux
}

// NewServer creates a new server.
func NewServer() (server *Server, err error) {
	server = &Server{}
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

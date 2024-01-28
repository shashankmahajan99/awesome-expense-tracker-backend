// Package apipkg is a package that contains the server and client code for the API.
package apipkg

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	AwesomeExpenseTrackerApi "github.com/shashankmahajan99/awesome-expense-tracker-backend/api"
	"github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/api/auth"
	db "github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/db/sqlc"
)

// Server is the gRPC server.
type Server struct {
	authenticationManager *auth.AuthenticationManager
	AwesomeExpenseTrackerApi.UnimplementedUserAuthenticationServer
	AwesomeExpenseTrackerApi.UnimplementedUserProfileServer
	AwesomeExpenseTrackerApi.UnimplementedExpenseManagementServer
	router *runtime.ServeMux
	store  db.Store
}

// NewServer creates a new server.
func NewServer(store db.Store, authenticationManager *auth.AuthenticationManager) *Server {
	return &Server{store: store, authenticationManager: authenticationManager}
}

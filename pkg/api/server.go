// Package apipkg is a package that contains the server and client code for the API.
package apipkg

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	AwesomeExpenseTrackerApi "github.com/shashankmahajan99/awesome-expense-tracker-backend/api"
	db "github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/db/sqlc"
	"golang.org/x/oauth2"
)

// Server is the gRPC server.
type Server struct {
	config *Config
	AwesomeExpenseTrackerApi.UnimplementedUserAuthenticationServer
	AwesomeExpenseTrackerApi.UnimplementedUserProfileServer
	router *runtime.ServeMux
	store  db.Store
}

// Config is the server config.
type Config struct {
	JwtKey         []byte
	GcpOAuthConfig *oauth2.Config
}

// NewServer creates a new server.
func NewServer(store db.Store, config *Config) (server *Server, err error) {
	server = &Server{store: store}

	server.config = config
	if err != nil {
		return nil, err
	}
	return server, nil
}

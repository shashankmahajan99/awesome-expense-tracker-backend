// `userauth` package defines the server API for UserAuthentication service.
package userauth

import (
	"context"

	"github.com/shashankmahajan99/awesome-expense-tracker-backend/api"
)

// UserAuthenticationServer is the server API for UserAuthentication service.
type UserAuthenticationServer interface {
	// Login logs in a user.
	Login(context.Context, *api.LoginRequest) (*api.LoginResponse, error)
	// Register registers a user.
	Register(context.Context, *api.RegisterRequest) (*api.RegisterResponse, error)
}

func (s *Server) Login(ctx context.Context, req *api.LoginRequest) (res *api.LoginResponse, err error) {
	return res, nil
}

func (s *Server) Register(ctx context.Context, req *api.RegisterRequest) (res *api.RegisterResponse, err error) {
	return res, nil
}

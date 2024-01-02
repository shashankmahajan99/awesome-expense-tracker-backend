package apipkg

import (
	"context"

	AwesomeExpenseTrackerApi "github.com/shashankmahajan99/awesome-expense-tracker-backend/api"
)

// UserAuthServer is the server API for UserAuthentication service.
type UserAuthServer interface {
	// Login logs in a user.
	Login(context.Context, *AwesomeExpenseTrackerApi.LoginRequest) (*AwesomeExpenseTrackerApi.LoginResponse, error)
	// Register registers a user.
	Register(context.Context, *AwesomeExpenseTrackerApi.RegisterRequest) (*AwesomeExpenseTrackerApi.RegisterResponse, error)
}

func (s *Server) Login(ctx context.Context, req *AwesomeExpenseTrackerApi.LoginRequest) (res *AwesomeExpenseTrackerApi.LoginResponse, err error) {
	// Add user login logic

	res = &AwesomeExpenseTrackerApi.LoginResponse{Token: "token"}
	return res, nil
}

func (s *Server) Register(ctx context.Context, req *AwesomeExpenseTrackerApi.RegisterRequest) (res *AwesomeExpenseTrackerApi.RegisterResponse, err error) {
	return res, nil
}

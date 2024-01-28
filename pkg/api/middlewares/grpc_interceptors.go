// Package middlewares is a package that contains the middlewares for the API.
package middlewares

import (
	"context"
	"strings"

	"github.com/bufbuild/protovalidate-go"
	"github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/api/auth"
	"github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// AuthInterceptor is a interceptor that checks if the user is authenticated.
type AuthInterceptor struct {
	authenticationManager *auth.AuthenticationManager
	accessibleRoles       map[string][]string
}

// NewAuthInterceptor creates a new AuthInterceptor.
func NewAuthInterceptor(authenticationManager *auth.AuthenticationManager, accessibleRoles map[string][]string) *AuthInterceptor {
	return &AuthInterceptor{authenticationManager: authenticationManager, accessibleRoles: accessibleRoles}
}

// GrpcAuthInterceptor is a interceptor that checks if the user is authenticated.
func (authInterceptor *AuthInterceptor) GrpcAuthInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Check if the method is accessible without authentication
		accessibleRoles, ok := authInterceptor.accessibleRoles[info.FullMethod]
		if !ok {
			// Public path accessible without authentication
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "failed to get metadata")
		}

		authorization := md.Get("authorization")
		if len(authorization) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "authorization header is required")
		}

		bearerString := strings.Split(authorization[0], " ")[0]
		if bearerString != "Bearer" {
			return nil, status.Errorf(codes.Unauthenticated, "authorization header is invalid")
		}

		encodedToken := strings.Split(authorization[0], " ")[1]
		if encodedToken == "" {
			return nil, status.Errorf(codes.Unauthenticated, "token is required")
		}

		userClaims, err := authInterceptor.authenticationManager.ValidateToken(ctx, encodedToken)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "invalid token: "+err.Error())
		}

		ctx = context.WithValue(ctx, utils.EmailKey, userClaims.Email)

		for _, role := range accessibleRoles {
			if role == userClaims.Role {
				return handler(ctx, req)
			}
		}

		return nil, status.Errorf(codes.PermissionDenied, "user doesn't have permission to access this resource")
	}
}

// GrpcRequestValidator is a interceptor that validates the request.
func GrpcRequestValidator(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to initialize validator: "+err.Error())
	}

	// Type assert req to the appropriate type that implements protoreflect.ProtoMessage
	message, ok := req.(protoreflect.ProtoMessage)
	if !ok {
		return nil, status.Error(codes.Internal, "failed to validate request: invalid request type")
	}

	err = v.Validate(message)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "failed to validate request: "+err.Error())
	}

	return handler(ctx, req)
}

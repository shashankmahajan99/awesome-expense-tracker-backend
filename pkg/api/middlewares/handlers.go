package middlewares

import (
	"log"
	"mime"
	"net/http"
	"strings"

	"github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/api/auth"
	failuremanagement "github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/failure_management"
	"github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/utils"
)

// AuthMiddleware is a middleware that checks if the user is authenticated.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var customError *failuremanagement.CustomErrorResponse
		w.Header().Set("Content-Type", "application/json")

		contentType := r.Header.Get("Content-Type")

		ctx := r.Context()

		if contentType != "" {
			mimeType, _, err := mime.ParseMediaType(contentType)
			if err != nil {
				customError = failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, "Malformed Content-Type header", http.StatusBadRequest)
				failuremanagement.NewHTTPCustomErrorResponse(w, customError)
				return
			}

			if mimeType != "application/json" {
				customError = failuremanagement.NewCustomErrorResponse(utils.INVALIDREQUEST, "Content-Type header must be application/json", http.StatusBadRequest)
				failuremanagement.NewHTTPCustomErrorResponse(w, customError)
				return
			}
		}

		// Get Authorization Header
		authorization := r.Header.Get("Authorization")

		if authorization == "" {
			customError = failuremanagement.NewCustomErrorResponse(utils.UNAUTHORIZEDREQUEST, "Authorization header is required", http.StatusUnauthorized)
			failuremanagement.NewHTTPCustomErrorResponse(w, customError)
			return
		}

		// Check only one space exists in authorization header
		bearerString := strings.Split(authorization, " ")[0]
		if bearerString != "Bearer" {
			customError = failuremanagement.NewCustomErrorResponse(utils.UNAUTHORIZEDREQUEST, "Authorization header is invalid", http.StatusUnauthorized)
			failuremanagement.NewHTTPCustomErrorResponse(w, customError)
			return
		}

		encodedToken := strings.Split(authorization, " ")[1]
		if encodedToken == "" {
			customError = failuremanagement.NewCustomErrorResponse(utils.UNAUTHORIZEDREQUEST, "Token is required", http.StatusUnauthorized)
			failuremanagement.NewHTTPCustomErrorResponse(w, customError)
			return
		}

		// Validate token
		tokenHandler := auth.NewTokenHandler()
		tokenClaims, err := tokenHandler.ValidateToken(ctx, nil, encodedToken)
		if err != nil {
			customError = failuremanagement.NewCustomErrorResponse(utils.UNAUTHORIZEDREQUEST, "Invalid token", http.StatusUnauthorized)
			failuremanagement.NewHTTPCustomErrorResponse(w, customError)
			return
		}

		log.Default().Println("User Authorized access: " + tokenClaims.Email)
		next.ServeHTTP(w, r)
	})
}

package auth

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	apipkg "github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/api"
	"google.golang.org/api/idtoken"
)

// TokenHandler is a handler for token requests.
type TokenHandler struct {
}

// NewTokenHandler creates a new token handler.
func NewTokenHandler() *TokenHandler {
	return &TokenHandler{}
}

type TokenClaims struct {
	Email string `json:"email"`
}

// ValidateToken validates a token.
func (h *TokenHandler) ValidateToken(ctx context.Context, config *apipkg.Config, encodedToken string) (verifiedClaims *TokenClaims, err error) {
	verifiedClaims = &TokenClaims{}
	unverifiedClaims, err := jwt.Parse(encodedToken, nil)
	// ignore if token signature is invalid or token is unverifiable
	if err != nil && !(strings.Contains(err.Error(), jwt.ErrSignatureInvalid.Error()) || strings.Contains(err.Error(), jwt.ErrTokenUnverifiable.Error())) {
		return nil, err
	}

	expirationTime, err := unverifiedClaims.Claims.GetExpirationTime()
	if err != nil {
		return nil, err
	}
	if time.Now().After(expirationTime.Time) {
		return nil, errors.New("token is expired")
	}

	issuer, err := unverifiedClaims.Claims.GetIssuer()
	if issuer == "awesome-expense-tracker" {
		verifyToken, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
			return config.JwtKey, nil
		})
		if err != nil {
			return nil, err
		}
		if !verifyToken.Valid {
			return nil, errors.New("token is invalid")
		}
		verifiedClaims.Email = unverifiedClaims.Claims.(jwt.MapClaims)["email"].(string)
	} else {
		audience, err := unverifiedClaims.Claims.GetAudience()
		if err != nil {
			return nil, err
		}

		validPaylod, err := idtoken.Validate(ctx, encodedToken, audience[0])
		if err != nil {
			return nil, err
		}
		verifiedClaims.Email = validPaylod.Claims["email"].(string)
	}

	return verifiedClaims, nil
}

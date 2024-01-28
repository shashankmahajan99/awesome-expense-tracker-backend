// Package auth is a package that contains authentication and authorization related code.
package auth

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/api/idtoken"
)

// ValidateToken validates a token.
func (authenticationManager *AuthenticationManager) ValidateToken(ctx context.Context, encodedToken string) (verifiedClaims *UserClaims, err error) {
	verifiedClaims = &UserClaims{}
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
	if err != nil {
		return nil, err
	}
	if issuer == "awesome-expense-tracker" {
		verifyToken, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
			return authenticationManager.jwtKey, nil
		})
		if err != nil {
			return nil, err
		}
		if !verifyToken.Valid {
			return nil, errors.New("token is invalid")
		}
		verifiedClaims.Email = unverifiedClaims.Claims.(jwt.MapClaims)["email"].(string)
		verifiedClaims.Role = unverifiedClaims.Claims.(jwt.MapClaims)["role"].(string)
		verifiedClaims.Issuer = issuer
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
		verifiedClaims.Role = "admin"
		verifiedClaims.Issuer = issuer
	}

	return verifiedClaims, nil
}

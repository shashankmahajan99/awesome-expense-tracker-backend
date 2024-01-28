package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
)

// AuthenticationManager is the auth manager.
type AuthenticationManager struct {
	jwtKey         []byte
	GcpOAuthConfig *oauth2.Config
}

// UserClaims is the user claims for the JWT token.
type UserClaims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
	Role  string `json:"role"`
}

// NewAuthenticationManager creates a new config manager.
func NewAuthenticationManager(jwtKey []byte, gcpOauthConfig *oauth2.Config) *AuthenticationManager {
	return &AuthenticationManager{
		jwtKey:         jwtKey,
		GcpOAuthConfig: gcpOauthConfig,
	}
}

// GenerateJWTToken generates a JWT token.
func (manager *AuthenticationManager) GenerateJWTToken(email string, role string) (*oauth2.Token, error) {
	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "awesome-expense-tracker",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
		Email: email,
		Role:  role,
	}

	idToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	idTokenString, err := idToken.SignedString(manager.jwtKey)
	if err != nil {
		return nil, err
	}
	token := &oauth2.Token{
		AccessToken:  "",
		RefreshToken: "",
		Expiry:       time.Now().Add(time.Hour * 1),
		TokenType:    "Bearer",
	}
	extra := make(map[string]interface{})
	extra["id_token"] = idTokenString
	token = token.WithExtra(extra)

	return token, nil
}

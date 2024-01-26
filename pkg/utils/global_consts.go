package utils

const (
	// GoogleUserInfoURL is the URL for retrieving user information from Google.
	GoogleUserInfoURL = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

	// Auth_providers
	GoogleAuthProvider = "google"
	CustomAuthProvider = "custom"

	// Error messages

	INVALIDREQUEST      = "invalid_request"
	INTERNALERROR       = "internal_error"
	UNAUTHORIZEDREQUEST = "unauthorized_request"
)

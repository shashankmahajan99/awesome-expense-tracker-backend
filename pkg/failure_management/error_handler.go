package failuremanagement

import (
	"encoding/json"
	"net/http"
)

// CustomErrorResponse is a custom error response.
type CustomErrorResponse struct {
	Err              string `json:"error"`
	ErrorDescription string `json:"error_description"`
	Code             int    `json:"code"`
}

// NewCustomErrorResponse creates a new custom error response.
func NewCustomErrorResponse(err string, errDesc string, code int) *CustomErrorResponse {
	return &CustomErrorResponse{
		Err:              err,
		ErrorDescription: errDesc,
		Code:             code,
	}
}

func (e *CustomErrorResponse) Error() string {
	errBytes, err := json.Marshal(e)
	if err != nil {
		return "{'error':'SOMETHING TERRIBLE WAS SENT TO THE CLIENT'}"
	}
	return string(errBytes)
}

func NewHTTPCustomErrorResponse(w http.ResponseWriter, customError *CustomErrorResponse) {
	w.WriteHeader(customError.Code)
	w.Write([]byte(customError.Error()))
}

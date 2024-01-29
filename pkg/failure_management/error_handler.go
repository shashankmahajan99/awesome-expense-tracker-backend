package failuremanagement

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/status"
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

// CustomHTTPErrorWrapper is a custom error wrapper.
func CustomHTTPErrorWrapper(_ context.Context, _ *runtime.ServeMux, _ runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json")
	CustomHTTPError(w, err)
}

// CustomHTTPError is a custom error handler.
func CustomHTTPError(w http.ResponseWriter, err error) {
	statusObject := status.Convert(err)
	errorMessage := statusObject.Message()

	// Retain HTTP error codes
	var httpCode int
	switch statusObject.Code() {
	case 0:
		httpCode = http.StatusOK
	case 1:
		httpCode = http.StatusRequestTimeout
	case 2:
		httpCode = http.StatusInternalServerError
	case 3:
		httpCode = http.StatusBadRequest
	case 4:
		httpCode = http.StatusGatewayTimeout
	case 5:
		httpCode = http.StatusNotFound
	case 6:
		httpCode = http.StatusConflict
	case 7:
		httpCode = http.StatusForbidden
	case 16:
		httpCode = http.StatusUnauthorized
	case 8:
		httpCode = http.StatusTooManyRequests
	case 9:
		httpCode = http.StatusBadRequest
	case 10:
		httpCode = http.StatusConflict
	case 11:
		httpCode = http.StatusBadRequest
	case 12:
		httpCode = http.StatusNotImplemented
	case 13:
		httpCode = http.StatusInternalServerError
	case 14:
		httpCode = http.StatusServiceUnavailable
	case 15:
		httpCode = http.StatusInternalServerError
	default:
		httpCode = http.StatusInternalServerError
	}

	w.WriteHeader(httpCode)

	customError := &CustomErrorResponse{}

	// Unmarshal the error string into the CustomErrorResponse struct
	jsonErr := json.Unmarshal([]byte(errorMessage), customError)
	if jsonErr == nil {
		w.Write([]byte(customError.Error()))
		return
	}

	// If the error is a grpc Error, then write the grpc error message
	err = NewCustomErrorResponse("rpc error", errorMessage, httpCode)
	w.Write([]byte(err.Error()))
	return
}

package api

import (
	"encoding/json"
)

type ErrorResponse struct {
	Error   string  `json:"error"`
	Message *string `json:"message,omitempty"`
}

func newErrorResponse(errorType string, errorMessage *string) ([]byte, error) {
	errorResponse := &ErrorResponse{
		Error:   errorType,
		Message: errorMessage,
	}

	return json.Marshal(errorResponse)
}

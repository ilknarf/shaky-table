package api

import (
	"encoding/json"
)

type CreateAccountResponse struct {
	Success bool    `json:"success"`
	Message *string `json:"message,omitempty"`
}

func newCreateAccountResponse(isError bool, message *string) []byte {
	response := &CreateAccountResponse{
		Success: true,
	}

	if isError {
		response.Success = false
		response.Message = message
	}

	resp, _ := json.Marshal(response)

	return resp
}

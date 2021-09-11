package api

import (
	"encoding/json"
)

type Response struct {
	Success bool    `json:"success"`
	Message *string `json:"message,omitempty"`
}

func newResponse(isError bool, message *string) []byte {
	response := &Response{
		Success: true,
	}

	if isError {
		response.Success = false
		response.Message = message
	}

	resp, _ := json.Marshal(response)

	return resp
}

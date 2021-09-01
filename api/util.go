package api

import (
	"log"
	"mime"
	"net/http"
)

func validPOST(w http.ResponseWriter, r *http.Request, apiName string) (bool, int) {
	if r.Method != "POST" {
		log.Printf("Incorrect request type for %s (expected POST)\n", apiName)
		return false, http.StatusMethodNotAllowed
	}

	contentType := r.Header.Get("Content-type")
	if contentType == "" {
		log.Printf("Missing Content-type header for %s\n", apiName)
		return false, http.StatusNoContent
	}

	if mediaType, _, err := mime.ParseMediaType(contentType); err != nil || mediaType != "application/x-www-form-urlencoded" {
		log.Printf("Incorrect Content-type header for %s\n", apiName)
		return false, http.StatusUnsupportedMediaType
	}

	return true, 0
}

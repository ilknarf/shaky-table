package api

import (
	"errors"
	"mime"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func validPOST(r *http.Request) (error, int) {
	if r.Method != "POST" {
		err := errors.New("Incorrect request type for request (expected POST)")
		return err, http.StatusMethodNotAllowed
	}

	contentType := r.Header.Get("Content-type")
	if contentType == "" {
		err := errors.New("Missing Content-type header for %s\n")
		return err, http.StatusNoContent
	}

	if mediaType, _, err := mime.ParseMediaType(contentType); err != nil || mediaType != "application/x-www-form-urlencoded" {
		err := errors.New("Incorrect Content-type header for %s\n")
		return err, http.StatusUnsupportedMediaType
	}

	return nil, 0
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

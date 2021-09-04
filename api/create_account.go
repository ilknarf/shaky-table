package api

import (
	"log"
	"net/http"

	"github.com/pkg/errors"
)

func (api *API) CreateAccount(w http.ResponseWriter, r *http.Request) {
	if err, responseCode := validPOST(r); err != nil {
		log.Println(errors.Wrap(err, "CreateAccount error"))
		w.WriteHeader(responseCode)
		return
	}

	ctx := r.Context()
	username := r.Form.Get("user")
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	if username == "" || password == "" {
		log.Println(errors.New("CreateAccount Error: missing signup form input"))
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")

		errorType := "bad input"
		errorMessage := "Missing form inputs for signup. Please try again with all required fields"
		errorResponse, err := newErrorResponse(errorType, &errorMessage)
		if err != nil {
			log.Println(errors.Wrap(err, "Unable to create errorResponse"))
			return
		}

		w.Write(errorResponse)

		return
	}

	if api.userDB.UserExists(ctx, username) {
		w.WriteHeader(http.StatusConflict)
		return
	}

	if err := api.userDB.CreateUser(ctx, username, password, email); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

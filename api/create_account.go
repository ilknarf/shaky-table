package api

import (
	"log"
	"net/http"
)

func (api *API) CreateAccount(w http.ResponseWriter, r *http.Request) {
	if err, responseCode := validPOST(r); err != nil {
		log.Print(err)
		w.WriteHeader(responseCode)
		return
	}

	ctx := r.Context()
	username := r.Form.Get("user")
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	if username == "" || password == "" {
		w.WriteHeader(http.StatusBadRequest)
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

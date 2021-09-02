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

	username := r.Form.Get("user")
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	if username == "" || password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := api.userDB.CreateUser(username, password, email); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

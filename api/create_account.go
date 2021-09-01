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

	if err := api.userDB.CreateUser("user", "pass"); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

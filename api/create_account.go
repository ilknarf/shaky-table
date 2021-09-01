package api

import (
	"net/http"
)

func (api *API) CreateAccount(w http.ResponseWriter, r *http.Request) {
	ok, responseCode := validPOST(w, r, "CreateAccount")

	if !ok {
		w.WriteHeader(responseCode)
		return
	}

}

package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

func (api *API) CreateAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	username := r.Form.Get("username")
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	var (
		responseCode int
		response     []byte
	)

	w.Header().Set("Content-Type", "application/json")

	// invalid request body/formatting
	if err, respCode := validPOST(r); err != nil {
		log.Println(errors.Wrap(err, "CreateAccount error: incorrect body type"))
		responseCode = respCode

		errorMessage := "invalid request type"
		response = newCreateAccountResponse(true, &errorMessage)
		// missing user/pass field
	} else if username == "" || password == "" {
		log.Println(errors.New("CreateAccount: missing signup form input on attempt"))
		responseCode = http.StatusBadRequest

		missing := make([]string, 0)
		if username == "" {
			missing = append(missing, "username")
		}

		if password == "" {
			missing = append(missing, "password")
		}

		errorMessage := "Missing (" + strings.Join(missing, ", ") + ") for signup. Please try again with all required fields"
		response = newCreateAccountResponse(true, &errorMessage)
	} else if api.userDB.UserExists(ctx, username) {
		log.Println("CreateAccount duplicate username signup attempted")
		responseCode = http.StatusConflict

		errorMessage := "user already exists"
		response = newCreateAccountResponse(true, &errorMessage)

	} else if err := api.userDB.CreateUser(ctx, username, password, email); err != nil {
		log.Println(err)
		responseCode = http.StatusInternalServerError

		errorMessage := "unable to create new user"
		response = newCreateAccountResponse(true, &errorMessage)
	} else {
		// success
		log.Println(fmt.Sprintf("CreateAccount new username %s", username))

		responseCode = http.StatusOK
		response = newCreateAccountResponse(false, nil)
	}

	w.WriteHeader(responseCode)
	w.Write(response)
}

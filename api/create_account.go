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

	var (
		responseCode int
		response     []byte
	)

	setErrorResponse := func(respCode int, message string) {
		responseCode = respCode
		response = newCreateAccountResponse(true, &message)
	}

	w.Header().Set("Content-Type", "application/json")

	// invalid request body/formatting
	if err, respCode := validPOST(r); err != nil {
		log.Println(errors.Wrap(err, "CreateAccount incorrect body type"))
		setErrorResponse(respCode, "invalid request type")

		w.WriteHeader(responseCode)
		w.Write(response)
		return
	}

	if err := r.ParseForm(); err != nil {
		log.Println(errors.Wrap(err, "CreateAccount unable to parse form data"))
		setErrorResponse(http.StatusBadRequest, "unable to parse form data")

		w.WriteHeader(responseCode)
		w.Write(response)
		return
	}

	username := r.Form.Get("username")
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	if username == "" || password == "" {
		log.Println(errors.New("CreateAccount missing signup form input on attempt"))

		missing := make([]string, 0)
		if username == "" {
			missing = append(missing, "username")
		}

		if password == "" {
			missing = append(missing, "password")
		}

		errorMessage := "Missing (" + strings.Join(missing, ", ") + ") for signup. Please try again with all required fields"
		setErrorResponse(http.StatusBadRequest, errorMessage)
	} else if exists, err := api.userDB.UserExists(ctx, username); exists || err != nil {
		if exists {
			log.Println("CreateAccount duplicate username signup attempted")
			setErrorResponse(http.StatusConflict, "user already exists")
		} else {
			log.Println(errors.Wrap(err, "CreateAccount unable to check if user exists"))
			setErrorResponse(http.StatusInternalServerError, "server was unable to handle the request")
		}
	} else if err := api.userDB.CreateUser(ctx, username, password, email); err != nil {
		log.Println(errors.Wrap(err, "CreateAccount unable to create user"))
		setErrorResponse(http.StatusInternalServerError, "unable to create new user")
	} else {
		// success
		log.Println(fmt.Sprintf("CreateAccount new username %s", username))

		responseCode = http.StatusOK
		response = newCreateAccountResponse(false, nil)
	}

	w.WriteHeader(responseCode)
	w.Write(response)
}

package api

import (
	"log"
	"net/http"
	"strings"

	"github.com/ilknarf/shaky-table/auth"
	"github.com/pkg/errors"
)

func (api *API) LoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var (
		responseCode int
		response     []byte
	)

	setErrorResponse := func(respCode int, message string) {
		responseCode = respCode
		response = newResponse(true, &message)
	}

	w.Header().Set("Content-Type", "application/json")

	// invalid request body/formatting
	if err, respCode := validPOST(r); err != nil {
		log.Println(errors.Wrap(err, "LoginUser incorrect body type"))
		setErrorResponse(respCode, "invalid request type")

		w.WriteHeader(responseCode)
		w.Write(response)
		return
	}

	if err := r.ParseForm(); err != nil {
		log.Println(errors.Wrap(err, "LoginUser unable to parse form data"))
		setErrorResponse(http.StatusBadRequest, "unable to parse form data")

		w.WriteHeader(responseCode)
		w.Write(response)
		return
	}

	username := r.Form.Get("username")
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

		errorMessage := "Missing (" + strings.Join(missing, ", ") + ") for login. Please try again with all required fields"
		setErrorResponse(http.StatusBadRequest, errorMessage)
	} else {
		u, err := api.userDB.VerifyLogin(ctx, username, password)
		if err != nil {
			errorMessage := "Unable to login. Please try again with the correct credentials"
			setErrorResponse(http.StatusBadRequest, errorMessage)
		} else {
			// map fields from userdb.User to auth.User
			api.authentication.AddUserToSession(w, r, auth.User{
				Username: u.Username,
				IsAdmin:  u.IsAdmin,
				LoggedIn: true,
			})

			w.WriteHeader(http.StatusOK)
			response = newResponse(false, nil)
		}
	}

	w.WriteHeader(responseCode)
	w.Write(response)
}

package api

import (
	"github.com/ilknarf/shaky-table/auth"
	"github.com/ilknarf/shaky-table/userdb"
)

// API contains relevant db connections for handlers, and contains the handlers as methods
type API struct {
	userDB         *userdb.UserDB
	authentication *auth.Auth
}

func newAPI(userDB *userdb.UserDB, authentication *auth.Auth) *API {
	return &API{
		userDB:         userDB,
		authentication: authentication,
	}
}

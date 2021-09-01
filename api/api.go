package api

import (
	"github.com/ilknarf/shaky-table/userdb"
)

// API contains relevant db connections for handlers, and contains the handlers as methods
type API struct {
	userDB *userdb.UserDB
}

func newAPI(userDB *userdb.UserDB) *API {
	return &API{
		userDB: userDB,
	}
}

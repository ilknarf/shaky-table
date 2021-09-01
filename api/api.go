package api

import (
	"database/sql"
)

// API contains relevant db connections for handlers, and contains the handlers as methods
type API struct {
	db *sql.DB
}

func newAPI(db *sql.DB) *API {
	return &API{
		db: db,
	}
}

package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

type APIRouter struct {
	*mux.Router
}

func NewRouter(db *sql.DB) http.Handler {
	api := newAPI(db)

	r := &APIRouter{mux.NewRouter()}
	r.registerHandlers(api)

	return r
}

func (r *APIRouter) registerHandlers(api *API) {
	post := r.Methods("POST")
	// get := r.Methods("GET")

	post.HandlerFunc(api.CreateAccount)
}

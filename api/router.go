package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ilknarf/shaky-table/userdb"
)

type APIRouter struct {
	router *mux.Router
}

func (r *APIRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), 500*time.Millisecond)
	defer cancel()
	r.router.ServeHTTP(w, req.WithContext(ctx))
}

// NewRouter creates the API struct that contains all the handler methods
func NewRouter(userDB *userdb.UserDB) http.Handler {
	api := newAPI(userDB)

	r := &APIRouter{mux.NewRouter()}
	r.registerHandlers(api)

	return r
}

func (r *APIRouter) registerHandlers(api *API) {
	subRouter := r.router.PathPrefix("/api").Subrouter()
	post := subRouter.Methods("POST").Subrouter()
	// get := subRouter.Methods("GET")

	post.HandleFunc("/create_account", api.CreateAccount)
}

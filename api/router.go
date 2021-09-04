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

// AddAPIRoutes creates the API struct that contains all the handler methods,
// and adds the handlers to the router
func AddAPIRoutes(userDB *userdb.UserDB, router *mux.Router) http.Handler {
	api := newAPI(userDB)

	r := &APIRouter{router}
	r.registerHandlers(api)

	return r
}

func (r *APIRouter) registerHandlers(api *API) {
	post := r.router.Methods("POST").Subrouter()
	// get := r.router.Methods("GET")

	post.HandleFunc("/create_account", api.CreateAccount)
}

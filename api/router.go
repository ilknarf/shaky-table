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
	ctx, _ := context.WithTimeout(req.Context(), 500*time.Millisecond)
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
	post := r.router.Methods("POST")
	// get := r.router.Methods("GET")

	post.HandlerFunc(api.CreateAccount)
}

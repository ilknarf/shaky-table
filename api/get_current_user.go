package api

import (
	"net/http"

	"github.com/ilknarf/shaky-table/auth"
)

func (api *API) GetCurrentUser(r *http.Request) (*auth.User, error) {
	return api.authentication.GetUser(r)
}

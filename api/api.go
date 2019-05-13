package api

import (
	"github.com/gorilla/mux"
	"github.com/gusta-project/go-api/model"
)

// API so we can use the Manager aka the DB in the Handlers
type API struct {
	model *model.Manager
}

// FIXME: split this in flavor & vendor API?

// New API
func New(m *model.Manager) *API {
	return &API{model: m}
}

// Register routes
func (a *API) Register(r *mux.Router) {
	a.VendorAPI().Register(r)
	a.FlavorAPI().Register(r)
}

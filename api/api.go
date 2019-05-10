package api

import (
	"github.com/gorilla/mux"
	"github.com/pscn/flavor2go/model"
)

type API struct {
	m *model.Manager
}

func New(m *model.Manager) *API {
	return &API{m: m}
}

// Register routes
func (a *API) Register(r *mux.Router) {
	r.HandleFunc("/vendor/", a.CreateVendor).Methods("POST")
	r.HandleFunc("/vendor/", a.UpdateVendor).Methods("PUT")
	r.HandleFunc("/vendor/{uuid}", a.UpdateVendor).Methods("PUT")
	r.HandleFunc("/vendor/{uuid}", a.GetVendorByUUID).Methods("GET")
	r.HandleFunc("/vendors/", a.GetVendors).Methods("GET")

	r.HandleFunc("/flavor/", a.CreateFlavor).Methods("POST")
	r.HandleFunc("/flavor/", a.UpdateFlavor).Methods("PUT")
	r.HandleFunc("/flavor/{uuid}", a.UpdateFlavor).Methods("PUT")
	r.HandleFunc("/flavor/{uuid}", a.GetFlavorByUUID).Methods("GET")
	r.HandleFunc("/flavors/", a.GetFlavors).Methods("GET")

}

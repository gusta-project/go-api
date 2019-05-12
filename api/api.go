package api

import (
	"github.com/gorilla/mux"
	"github.com/pscn/flavor2go/model"
)

// API so we can use the Manager aka the DB in the Handlers
type API struct {
	db     *model.Manager
	vendor *model.VendorManager
	flavor *model.FlavorManager
}

// FIXME: split this in flavor & vendor API?

// New API
func New(m *model.Manager) *API {
	return &API{
		db:     m,
		vendor: m.Vendor,
		flavor: m.Flavor,
	}
}

// Register routes
func (a *API) Register(r *mux.Router) {
	r.HandleFunc("/vendor/", a.AddVendor).Methods("POST")
	// update with slug as payload
	r.HandleFunc("/vendor/", a.UpdateVendor).Methods("PUT")
	// update with slug as payload
	r.HandleFunc("/vendor/{slug}", a.UpdateVendor).Methods("PUT")

	r.HandleFunc("/vendor/{slug}", a.GetVendor).Methods("GET")

	r.HandleFunc("/vendor/", a.DeleteVendor).Methods("DELETE")
	r.HandleFunc("/vendor/{slug}", a.DeleteVendor).Methods("DELETE")

	// FIXME: only for development
	r.HandleFunc("/vendors/", a.GetVendors).Methods("GET")

	r.HandleFunc("/flavor/", a.CreateFlavor).Methods("POST")
	// update with slug as payload
	r.HandleFunc("/flavor/", a.UpdateFlavor).Methods("PUT")
	// update with uuid in the URL
	r.HandleFunc("/flavor/{slug}", a.UpdateFlavor).Methods("PUT")

	// get with slug as payload
	r.HandleFunc("/flavor/", a.GetFlavor).Methods("GET")
	// get from slug
	r.HandleFunc("/flavor/{slug}", a.GetFlavor).Methods("GET")
	// FIXME: only for development
	r.HandleFunc("/flavors/", a.GetFlavors).Methods("GET")

}

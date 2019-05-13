package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gusta-project/go-api/model"
)

// VendorAPI -
type VendorAPI struct {
	model *model.VendorManager
}

// VendorAPI -
func (a *API) VendorAPI() *VendorAPI {
	return &VendorAPI{model: a.model.Vendor}
}

// Register routes for vendor
func (a *VendorAPI) Register(r *mux.Router) {
	r.HandleFunc("/vendor/", a.Create).Methods("POST")
	// update with slug as payload
	r.HandleFunc("/vendor/", a.Update).Methods("PUT")
	// update with slug as payload
	r.HandleFunc("/vendor/{slug}", a.Update).Methods("PUT")

	r.HandleFunc("/vendor/{slug}", a.Get).Methods("GET")

	r.HandleFunc("/vendor/", a.Delete).Methods("DELETE")
	r.HandleFunc("/vendor/{slug}", a.Delete).Methods("DELETE")

	// FIXME: only for development
	r.HandleFunc("/vendors/", a.GetAll).Methods("GET")
}

// Create -
func (a *VendorAPI) Create(w http.ResponseWriter, r *http.Request) {
	vendor := &model.Vendor{}
	err := json.NewDecoder(r.Body).Decode(vendor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Error(err))
		return
	}

	if err = a.model.Create(vendor); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Error(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vendor)
}

// Delete -
func (a *VendorAPI) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["slug"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Error(errors.New("slug missing")))
		return
	}
	if err := a.model.Delete(&model.Vendor{Slug: vars["slug"]}); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Error(err))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(Error(nil))
	return
}

// Update from JSON
func (a *VendorAPI) Update(w http.ResponseWriter, r *http.Request) {
	vendor := &model.Vendor{}
	err := json.NewDecoder(r.Body).Decode(vendor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Error(errors.New("failed to decode JSON payload")))
		return
	}
	vars := mux.Vars(r)
	// slug from URL > slug from body
	if vars["slug"] != "" {
		vendor.Slug = vars["slug"]
	}

	if err = a.model.Update(vendor); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Error(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vendor)
}

// Get -
func (a *VendorAPI) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vendor := a.model.Get(&model.Vendor{Slug: vars["slug"]})
	if vendor == nil || vendor.Slug == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write(Error(errors.New("vendor not found")))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vendor)
}

// GetAll -
func (a *VendorAPI) GetAll(w http.ResponseWriter, r *http.Request) {
	// FIXME: page, order...
	vendors := a.model.GetAll()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vendors)
}

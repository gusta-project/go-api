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
	// CREATE
	r.HandleFunc("/vendor/", a.Create).Methods("POST")

	// READ
	r.HandleFunc("/vendor/{slug}", a.Get).Methods("GET")
	r.HandleFunc("/vendors/", a.GetAll).Methods("GET")

	// UPDATE
	r.HandleFunc("/vendor/{slug}", a.Update).Methods("PUT")

	// DELETE
	r.HandleFunc("/vendor/{slug}", a.Delete).Methods("DELETE")
}

// Create -
func (a *VendorAPI) Create(w http.ResponseWriter, r *http.Request) {
	vendor := &model.Vendor{}
	err := json.NewDecoder(r.Body).Decode(vendor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Error(errors.New("failed to decode JSON payload")))
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
	err := a.model.Delete(&model.Vendor{Slug: vars["slug"]})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Write(Error(err))
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
	if vars["slug"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Error(errors.New("slug missing")))
		return
	}

	err = a.model.Update(&model.Vendor{Slug: vars["slug"]}, vendor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Write(Error(err))
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

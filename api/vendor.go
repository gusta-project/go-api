package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pscn/flavor2go/model"
)

// AddVendor -
func (a *API) AddVendor(w http.ResponseWriter, r *http.Request) {
	vendor := &model.Vendor{}
	err := json.NewDecoder(r.Body).Decode(vendor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Error(err))
		return
	}

	if err = a.vendor.Add(vendor); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Error(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vendor)
}

// DeleteVendor -
func (a *API) DeleteVendor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["slug"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Error(errors.New("slug missing")))
		return
	}
	if err := a.vendor.Delete(&model.Vendor{Slug: vars["slug"]}); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Error(err))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(Error(nil))
	return
}

// UpdateVendor from JSON
func (a *API) UpdateVendor(w http.ResponseWriter, r *http.Request) {
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

	if err = a.vendor.Update(vendor); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Error(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vendor)
}

// GetVendor -
func (a *API) GetVendor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vendor := a.vendor.Get(&model.Vendor{Slug: vars["slug"]})
	if vendor == nil || vendor.Slug == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write(Error(errors.New("vendor not found")))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vendor)
}

// GetVendors -
func (a *API) GetVendors(w http.ResponseWriter, r *http.Request) {
	vendors := a.vendor.GetAll()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vendors)
}

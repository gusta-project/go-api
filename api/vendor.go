package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gusta-project/go-api/model"
)

// AddVendor -
func (a *API) AddVendor(w http.ResponseWriter, r *http.Request) {
	vendor := &model.Vendor{}
	err := json.NewDecoder(r.Body).Decode(vendor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = a.m.AddVendor(vendor); err != nil {
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
	if vars["uuid"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Error(errors.New("UUID missing")))
		return
	}
	if err := a.m.DeleteVendor(&model.Vendor{UUID: vars["uuid"]}); err != nil {
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
		return
	}
	vars := mux.Vars(r)
	// uuid from URL > uuid from body
	if vars["uuid"] != "" {
		vendor.UUID = vars["uuid"]
	}

	if err = a.m.UpdateVendor(vendor); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Error(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vendor)
}

// GetVendorByUUID -
func (a *API) GetVendorByUUID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vendor := a.m.GetVendor(&model.Vendor{UUID: vars["uuid"]})
	if vendor.UUID == "" {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(w).Encode(vendor)
}

// GetVendors -
func (a *API) GetVendors(w http.ResponseWriter, r *http.Request) {
	vendors := a.m.GetVendors()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vendors)
}

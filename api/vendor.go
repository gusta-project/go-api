package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pscn/flavor2go/model"
)

func (a *API) CreateVendor(w http.ResponseWriter, r *http.Request) {
	var vendor model.Vendor
	err := json.NewDecoder(r.Body).Decode(&vendor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	vendor = *a.m.AddVendor(&vendor)
	if vendor.UUID == "" {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(vendor)
}

func (a *API) UpdateVendor(w http.ResponseWriter, r *http.Request) {
	var vendor model.Vendor
	err := json.NewDecoder(r.Body).Decode(&vendor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	if vars["uuid"] != "" {
		vendor.UUID = vars["uuid"]
	}
	vendor = *a.m.UpdateVendor(&vendor)
	if vendor.UUID == "" {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(vendor)
}

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

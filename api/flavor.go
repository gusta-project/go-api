package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pscn/flavor2go/model"
)

// CreateFlavor -
func (a *API) CreateFlavor(w http.ResponseWriter, r *http.Request) {
	flavor := &model.Flavor{}
	err := json.NewDecoder(r.Body).Decode(flavor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = a.m.AddFlavor(flavor); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Error(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(flavor)
}

// UpdateFlavor -
func (a *API) UpdateFlavor(w http.ResponseWriter, r *http.Request) {
	flavor := &model.Flavor{}
	err := json.NewDecoder(r.Body).Decode(flavor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	// uuid from URL > uuid from body
	if vars["uuid"] != "" {
		flavor.UUID = vars["uuid"]
	}

	if err = a.m.UpdateFlavor(flavor); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Error(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(flavor)
}

// GetFlavor -
func (a *API) GetFlavor(w http.ResponseWriter, r *http.Request) {
	flavor := &model.Flavor{}
	err := json.NewDecoder(r.Body).Decode(flavor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	// slug from URL > uuid from body
	if vars["slug"] != "" {
		flavor.Slug = vars["slug"]
	}
	if flavor.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(w).Encode(flavor)
}

// GetFlavors -
func (a *API) GetFlavors(w http.ResponseWriter, r *http.Request) {
	flavors := a.m.GetFlavors()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(flavors)
}

package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gusta-project/go-api/model"
)

// FlavorAPI -
type FlavorAPI struct {
	model *model.FlavorManager
}

// FlavorAPI -
func (a *API) FlavorAPI() *FlavorAPI {
	return &FlavorAPI{model: a.model.Flavor}
}

// Register routes for Flavor
func (a *FlavorAPI) Register(r *mux.Router) {
	r.HandleFunc("/flavor/", a.Create).Methods("POST")
	// update with slug as payload
	r.HandleFunc("/flavor/", a.Update).Methods("PUT")
	// update with uuid in the URL
	r.HandleFunc("/flavor/{slug}", a.Update).Methods("PUT")

	// get with slug as payload
	r.HandleFunc("/flavor/", a.Get).Methods("GET")
	// get from slug
	r.HandleFunc("/flavor/{slug}", a.Get).Methods("GET")
	// FIXME: only for development
	r.HandleFunc("/flavors/", a.GetAll).Methods("GET")
}

// Create -
func (a *FlavorAPI) Create(w http.ResponseWriter, r *http.Request) {
	flavor := &model.Flavor{}
	err := json.NewDecoder(r.Body).Decode(flavor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Error(err))
		return
	}
	log.Printf("CreateFlavor: %#v", flavor)
	log.Printf("CreateFlavor: %#v", flavor.Vendor)

	if err = a.model.Create(flavor); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Error(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(flavor)
}

// Update -
func (a *FlavorAPI) Update(w http.ResponseWriter, r *http.Request) {
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

	if err = a.model.Update(flavor); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Error(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(flavor)
}

// Get -
func (a *FlavorAPI) Get(w http.ResponseWriter, r *http.Request) {
	flavor := &model.Flavor{}
	err := json.NewDecoder(r.Body).Decode(flavor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	// slug from URL > slug from body
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

// GetAll -
func (a *FlavorAPI) GetAll(w http.ResponseWriter, r *http.Request) {
	flavors := a.model.GetAll()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(flavors)
}

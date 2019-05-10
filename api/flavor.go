package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pscn/flavor2go/model"
)

func (a *API) GetFlavorByUUID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	flavor := a.m.GetFlavor(&model.Flavor{UUID: vars["uuid"]})
	if flavor.UUID == "" {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(w).Encode(flavor)
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/pscn/flavor2go/api"
	"github.com/pscn/flavor2go/model"

	"github.com/karrick/golf"
)

var port = golf.IntP('p', "port", 3000, "port to listen on")

func main() {
	golf.Parse()
	log.SetFlags(0)

	m := model.NewPostgres("localhost", 5432, "gusta", "gusta", "changeme", false)
	a := api.New(m)
	defer m.Close()
	r := mux.NewRouter()
	r.HandleFunc("/vendor/", a.CreateVendor).Methods("POST")
	r.HandleFunc("/vendor/", a.UpdateVendor).Methods("PUT")
	r.HandleFunc("/vendor/{uuid}", a.UpdateVendor).Methods("PUT")
	r.HandleFunc("/vendor/{uuid}", a.GetVendorByUUID).Methods("POST")
	r.HandleFunc("/flavor/{uuid}", a.GetFlavorByUUID).Methods("GET")

	addr := fmt.Sprintf(":%d", *port)
	log.Fatal(http.ListenAndServe(addr, r))

}

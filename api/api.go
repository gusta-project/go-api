package api

import "github.com/pscn/flavor2go/model"

type API struct {
	m *model.Manager
}

func New(m *model.Manager) *API {
	return &API{m: m}
}

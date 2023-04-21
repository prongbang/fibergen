package template

type wireApiTemplate struct {
	Module  string
	Project string
}

func (w *wireApiTemplate) Text() []byte {
	return []byte(`//go:build wireinject
// +build wireinject

package api

import (
	"github.com/google/wire"
	"` + w.Module + `/internal/` + w.Project + `/database"
	//+fibergen:import wire:package
)

func CreateAPI(dbDriver database.Drivers) API {
	wire.Build(
		NewAPI,
		NewRouters,
		//+fibergen:func wire:build
	)
	return nil
}`)
}

func WireApiTemplate(module string, project string) Template {
	return &wireApiTemplate{
		Module:  module,
		Project: project,
	}
}

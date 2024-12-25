package template

type wireTemplate struct {
	Module  string
	Project string
}

func (w *wireTemplate) Text() []byte {
	return []byte(`//go:build wireinject
// +build wireinject

package ` + w.Project + `

import (
	"github.com/google/wire"
	"` + w.Module + `/internal/` + w.Project + `"
	"` + w.Module + `/internal/` + w.Project + `/api"
	"` + w.Module + `/internal/` + w.Project + `/database"
	"` + w.Module + `/internal/pkg/response"
	"` + w.Module + `/internal/pkg/validator"
	//+fibergen:import wire:package
)

func CreateApp(dbDriver database.Drivers) ` + w.Project + `.App {
	wire.Build(
		` + w.Project + `.NewApp,
		api.New,
		api.NewRouters,
		response.New,
		validator.New,
		//+fibergen:func wire:build
	)
	return nil
}`)
}

func WireTemplate(module string, project string) Template {
	return &wireTemplate{
		Module:  module,
		Project: project,
	}
}

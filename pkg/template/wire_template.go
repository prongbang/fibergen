package template

import (
	"github.com/prongbang/fibergen/pkg/tocase"
)

type wireTemplate struct {
	Module  string
	Project string
}

func (w *wireTemplate) Text() []byte {
	project := tocase.ToLower(w.Project)
	return []byte(`//go:build wireinject
// +build wireinject

package ` + project + `

import (
	"github.com/google/wire"
	"` + w.Module + `/internal"
	"` + w.Module + `/internal/api"
	"` + w.Module + `/internal/database"
	"` + w.Module + `/internal/pkg/response"
	"` + w.Module + `/internal/pkg/validator"
	//+fibergen:import wire:package
)

func CreateApp(dbDriver database.Drivers) internal.App {
	wire.Build(
		internal.NewApp,
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

package template

type wireGenApiTemplate struct {
	Module  string
	Project string
}

func (w *wireGenApiTemplate) Text() []byte {
	return []byte(`// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//go:build !wireinject
// +build !wireinject

package api

import (
	"` + w.Module + `/internal/` + w.Project + `/database"
)

// Injectors from wire.go:

func CreateAPI(dbDriver database.Drivers) API {
	apiRouters := NewRouters()
	apiAPI := NewAPI(apiRouters)
	return apiAPI
}
`)
}

func WireGenApiTemplate(module string, project string) Template {
	return &wireGenApiTemplate{
		Module:  module,
		Project: project,
	}
}
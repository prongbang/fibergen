package template

type appTemplate struct {
	Project string
	Module  string
}

func (a *appTemplate) Text() []byte {
	return []byte(`package ` + a.Project + `

import (
	"` + a.Module + `/internal/` + a.Project + `/api"
)

type App interface {
	StartAPI()
}

type app struct {
	API  api.API
}

func (a *app) StartAPI() {
	a.API.Register()
}

func NewApp(apis api.API) App {
	return &app{
		API:  apis,
	}
}`)
}

func AppTemplate(module string, project string) Template {
	return &appTemplate{
		Module:  module,
		Project: project,
	}
}

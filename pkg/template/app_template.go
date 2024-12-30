package template

type appTemplate struct {
	Module string
}

func (a *appTemplate) Text() []byte {
	return []byte(`package internal

import (
	"` + a.Module + `/internal/api"
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

func AppTemplate(module string) Template {
	return &appTemplate{
		Module: module,
	}
}

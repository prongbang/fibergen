package template

type apiRoutersTemplate struct {
	Module string
}

func (r *apiRoutersTemplate) Text() []byte {
	return []byte(`package api

import (
	"github.com/gofiber/fiber/v2"
	"` + r.Module + `/pkg/core"
	//+fibergen:import routers:package
)

type Routers interface {
	core.Routers
}

type routers struct {
	//+fibergen:struct routers
}

func (r *routers) Initials(app *fiber.App) {
	//+fibergen:func initials
}

func NewRouters(
	//+fibergen:func new:routers
) Routers {
	return &routers{
		//+fibergen:return &routers
	}
}`)
}

func ApiRoutersTemplate(module string) Template {
	return &apiRoutersTemplate{
		Module: module,
	}
}

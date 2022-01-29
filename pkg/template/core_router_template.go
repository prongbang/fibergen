package template

type coreRouterTemplate struct {
}

func (c *coreRouterTemplate) Text() []byte {
	return []byte(`package core

import "github.com/gofiber/fiber/v2"

type Router interface {
	Initial(app *fiber.App)
}

type Routers interface {
	Initials(app *fiber.App)
}
`)
}

func CoreRouterTemplate() Template {
	return &coreRouterTemplate{}
}

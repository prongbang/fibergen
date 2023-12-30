package otp

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/mvp/test_test/pkg/core"
)

type Router interface {
	core.Router
}

type router struct {
	Handle   Handler
	Validate Validate
}

func (r *router) Initial(app *fiber.App) {
}

func NewRouter(handle Handler, validate Validate) Router {
	return &router{
		Handle:   handle,
		Validate: validate,
	}
}
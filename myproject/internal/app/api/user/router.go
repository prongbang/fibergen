package user

import "github.com/gofiber/fiber/v2"

type Router interface {
	core.Router
}

type router struct {
	Handle Handler
}

func (r *router) Initial(app *fiber.App) {
}

func NewRouter(handle Handler) Router {
	return &router{
		Handle: handle,
	}
}
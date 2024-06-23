package template

import "fmt"

func Router(name string, moduleName string) string {
	return fmt.Sprintf(`package %s

	import (
		"github.com/gofiber/fiber/v2"
		"%s/pkg/core"
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
	}`, name, moduleName)
}

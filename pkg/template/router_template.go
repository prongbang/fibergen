package template

import (
	"strings"

	"github.com/ettle/strcase"
)

func Router(name string, moduleName string) string {
	tmpl := `package {name}

import (
	"github.com/gofiber/fiber/v2"
	"{module}/pkg/core"
)

type Router interface {
	core.Router
}

type router struct {
	Handle   Handler
	Validate Validate
}

func (r *router) Initial(app *fiber.App) {
	//v1 := app.Group("/v1")
	//{
	//	v1.Post("/{route}/example", r.Validate.Example, r.Handle.Example)
	//}
}

func NewRouter(handle Handler, validate Validate) Router {
	return &router{
		Handle:   handle,
		Validate: validate,
	}
}`

	tmpl = strings.ReplaceAll(tmpl, "{module}", moduleName)
	tmpl = strings.ReplaceAll(tmpl, "{route}", strcase.ToKebab(name))
	tmpl = strings.ReplaceAll(tmpl, "{model}", strcase.ToPascal(name))
	tmpl = strings.ReplaceAll(tmpl, "{name}", strings.ToLower(name))

	return tmpl
}

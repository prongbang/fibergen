package template

import (
	"github.com/prongbang/fibergen/pkg/tocase"
	"strings"

	"github.com/ettle/strcase"
)

func RouterCrud(name string, moduleName string) string {
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
	v1 := app.Group("/v1")
	{
		v1.Post("/{route}/one", r.Validate.FindById, r.Handle.FindById)
		v1.Post("/{route}/many", r.Validate.FindList, r.Handle.FindList)
		v1.Post("/{route}/many/lite", r.Validate.FindLiteList, r.Handle.FindLiteList)
		v1.Post("/{route}/create", r.Validate.Create, r.Handle.Create)
		v1.Post("/{route}/update", r.Validate.Update, r.Handle.Update)
		v1.Post("/{route}/delete", r.Validate.Delete, r.Handle.Delete)
	}
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
	tmpl = strings.ReplaceAll(tmpl, "{name}", tocase.ToLower(name))

	return tmpl
}

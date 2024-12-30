package template

import (
	"github.com/prongbang/fibergen/pkg/tocase"
	"strings"

	"github.com/ettle/strcase"
)

func Handler(name string) string {
	tmpl := `package {name}

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/fibererror"
	"github.com/prongbang/goerror"
)

type Handler interface {
	Example(c *fiber.Ctx) error
}

type handler struct {
	Uc       UseCase
	Response fibererror.Response
}

// Example
// @Tags {tags}
// @Summary Example
// @Accept json
// @Produce json
// @Success 200 {object} core.Success{data={model}}
// @Failure 400 {object} core.Error
// @Failure 404 {object} core.Error
// @Security JWTAuth
// @Router /v1/{route}/example [post]
func (h *handler) Example(c *fiber.Ctx) error {
	return h.Response.With(c).Response(goerror.NewOK(nil))
}

func NewHandler(uc UseCase, response fibererror.Response) Handler {
	return &handler{
		Uc: 	  uc,
		Response: response,
	}
}`

	tmpl = strings.ReplaceAll(tmpl, "{tags}", strcase.ToSnake(name))
	tmpl = strings.ReplaceAll(tmpl, "{model}", strcase.ToPascal(name))
	tmpl = strings.ReplaceAll(tmpl, "{name}", tocase.ToLower(name))
	tmpl = strings.ReplaceAll(tmpl, "{route}", strcase.ToKebab(name))

	return tmpl
}

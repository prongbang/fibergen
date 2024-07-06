package template

import (
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
	FindById(c *fiber.Ctx) error
	FindList(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type handler struct {
	Uc       UseCase
	Response fibererror.Response
}

// FindById
// @Tags {name}
// @Summary Find a {name} by id
// @Accept json
// @Produce json
// @Param query body QueryOne true "query"
// @Success 200 {object} core.Success{data={model}}
// @Failure 400 {object} core.Error
// @Failure 404 {object} core.Error
// @Security JWTAuth
// @Router /v1/{route}/one [post]
func (h *handler) FindById(c *fiber.Ctx) error {
	return h.Response.With(c).Response(goerror.NewOK(nil))
}

// FindList
// @Tags {name}
// @Summary Find a list of {name}
// @Accept json
// @Produce json
// @Param query body QueryMany true "query"
// @Success 200 {object} core.Success{data=core.Paging{list=[]{model}}}
// @Failure 400 {object} core.Error
// @Security JWTAuth
// @Router /v1/{route}/many [post]
func (h *handler) FindList(c *fiber.Ctx) error {
	return h.Response.With(c).Response(goerror.NewOK(nil))
}

// Create
// @Tags {name}
// @Summary Create a {name}
// @Accept json
// @Produce json
// @Param {name} body Create{model} true "{name}"
// @Success 201 {object} core.Success{data={model}}
// @Failure 400 {object} core.Error
// @Security JWTAuth
// @Router /v1/{route}/create [post]
func (h *handler) Create(c *fiber.Ctx) error {
	return h.Response.With(c).Response(goerror.NewOK(nil))
}

// Update
// @Tags {name}
// @Summary Update a {name}
// @Accept json
// @Produce json
// @Param {name} body Update{model} true "{name}"
// @Success 200 {object} core.Success{data={model}}
// @Failure 400 {object} core.Error
// @Failure 404 {object} core.Error
// @Security JWTAuth
// @Router /v1/{route}/update [post]
func (h *handler) Update(c *fiber.Ctx) error {
	return h.Response.With(c).Response(goerror.NewOK(nil))
}

// Delete
// @Tags {name}
// @Summary Delete a {name} by id
// @Accept json
// @Produce json
// @Param {name} body Delete{model} true "{name}"
// @Success 200 {object} core.Success
// @Failure 400 {object} core.Error
// @Failure 404 {object} core.Error
// @Security JWTAuth
// @Router /v1/{route}/delete [post]
func (h *handler) Delete(c *fiber.Ctx) error {
	return h.Response.With(c).Response(goerror.NewOK(nil))
}

func NewHandler(uc UseCase, response fibererror.Response) Handler {
	return &handler{
		Uc: 	  uc,
		Response: response,
	}
}`

	tmpl = strings.ReplaceAll(tmpl, "{model}", strcase.ToPascal(name))
	tmpl = strings.ReplaceAll(tmpl, "{name}", strings.ToLower(name))
	tmpl = strings.ReplaceAll(tmpl, "{route}", strcase.ToKebab(name))

	return tmpl
}

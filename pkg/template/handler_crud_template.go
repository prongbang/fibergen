package template

import (
	"strings"

	"github.com/ettle/strcase"
)

func HandlerCrud(
	model string,
	module string,
	pk string,
) string {
	tmpl := `package {name}

import (
	"{module}/pkg/core"
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
	q := QueryOne{}
	_ = c.BodyParser(&q)

	if r := h.Uc.FindById(q.Id); r.Id {findOneCondition} {
		return h.Response.With(c).Response(goerror.NewOK(r))
	}

	return h.Response.With(c).Response(goerror.NewNotFound())
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
	q := QueryMany{}
	_ = c.BodyParser(&q)

	params := Params{
		QueryMany: q,
	}

	getCount := func() int64 { return h.Uc.Count(params) }

	getData := func(limit int64, offset int64) interface{} {
		params.Limit = limit
		params.Offset = offset
		return h.Uc.FindList(params)
	}
	
	r := core.Pagination(q.Page, q.Limit, getCount, getData)

	return h.Response.With(c).Response(goerror.NewOK(r))
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
	b := Create{model}{}
	_ = c.BodyParser(&b)

	d, err := h.Uc.Create(&b)
	if err != nil {
		return h.Response.With(c).Response(err)
	}

	return h.Response.With(c).Response(goerror.NewCreated(d))
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
	b := Update{model}{}
	_ = c.BodyParser(&b)

	if r := h.Uc.FindById(b.Id); r.Id == b.Id {
		d, err := h.Uc.Update(&b)
		if err != nil {
			return h.Response.With(c).Response(err)
		}
		
		return h.Response.With(c).Response(goerror.NewOK(d))
	}

	return h.Response.With(c).Response(goerror.NewNotFound())
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
	b := Delete{model}{}
	_ = c.BodyParser(&b)

	if r := h.Uc.FindById(b.Id); r.Id == b.Id {
		if err := h.Uc.Delete(b.Id); err != nil {
			return h.Response.With(c).Response(err)
		}
		
		return h.Response.With(c).Response(goerror.NewOK(nil))
	}

	return h.Response.With(c).Response(goerror.NewNotFound())
}

func NewHandler(uc UseCase, response fibererror.Response) Handler {
	return &handler{
		Uc: 	  uc,
		Response: response,
	}
}`

	if pk == "string" {
		tmpl = strings.ReplaceAll(tmpl, "{findOneCondition}", `!= ""`)
	} else if strings.Contains(pk, "float") || strings.Contains(pk, "int") {
		tmpl = strings.ReplaceAll(tmpl, "{findOneCondition}", `> 0`)
	} else {
		tmpl = strings.ReplaceAll(tmpl, "{findOneCondition}", `!= nil`)
	}

	tmpl = strings.ReplaceAll(tmpl, "{module}", module)
	tmpl = strings.ReplaceAll(tmpl, "{model}", strcase.ToPascal(model))
	tmpl = strings.ReplaceAll(tmpl, "{name}", strings.ToLower(model))
	tmpl = strings.ReplaceAll(tmpl, "{route}", strcase.ToKebab(model))

	return tmpl
}

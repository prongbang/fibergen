package template

import (
	"github.com/prongbang/fibergen/pkg/tocase"
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
	"{module}/pkg/requestx"
)

type Handler interface {
	FindById(c *fiber.Ctx) error
	FindList(c *fiber.Ctx) error
	FindLiteList(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type handler struct {
	Uc       UseCase
	Response fibererror.Response
}

// FindById
// @Tags {tags}
// @Summary Find a {tags} by id
// @Accept json
// @Produce json
// @Param query body QueryOne true "query"
// @Success 200 {object} core.Success{data={model}}
// @Failure 400 {object} core.Error
// @Failure 404 {object} core.Error
// @Security JWTAuth
// @Router /v1/{route}/one [post]
func (h *handler) FindById(c *fiber.Ctx) error {
	q := *requestx.Get[QueryOne](c)

	if r := h.Uc.FindById(q.Id); r.Id {findOneCondition} {
		return h.Response.With(c).Response(goerror.NewOK(r))
	}

	return h.Response.With(c).Response(goerror.NewNotFound())
}

// FindList
// @Tags {tags}
// @Summary Find a list of {tags}
// @Accept json
// @Produce json
// @Param query body QueryMany true "query"
// @Success 200 {object} core.Success{data=core.Paging{list=[]{model}}}
// @Failure 400 {object} core.Error
// @Security JWTAuth
// @Router /v1/{route}/many [post]
func (h *handler) FindList(c *fiber.Ctx) error {
	q := *requestx.Get[QueryMany](c)

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

// FindLiteList
// @Tags {tags}
// @Summary Find a lite list of {tags}
// @Accept json
// @Produce json
// @Param query body QueryMany true "query"
// @Success 200 {object} core.Success{data=[]Lite{model}}
// @Failure 400 {object} core.Error
// @Security JWTAuth
// @Router /v1/{route}/many/lite [post]
func (h *handler) FindLiteList(c *fiber.Ctx) error {
	q := *requestx.Get[LiteQueryMany](c)

	params := LiteParams{
		LiteQueryMany: q,
	}

	data := h.Uc.FindLiteList(params)

	return h.Response.With(c).Response(goerror.NewOK(data))
}

// Create
// @Tags {tags}
// @Summary Create a {tags}
// @Accept json
// @Produce json
// @Param query body Create{model} true "query"
// @Success 201 {object} core.Success{data={model}}
// @Failure 400 {object} core.Error
// @Security JWTAuth
// @Router /v1/{route}/create [post]
func (h *handler) Create(c *fiber.Ctx) error {
	b := *requestx.Get[Create{model}](c)

	d, err := h.Uc.Create(&b)
	if err != nil {
		return h.Response.With(c).Response(err)
	}

	return h.Response.With(c).Response(goerror.NewCreated(d))
}

// Update
// @Tags {tags}
// @Summary Update a {tags}
// @Accept json
// @Produce json
// @Param query body Update{model} true "query"
// @Success 200 {object} core.Success{data={model}}
// @Failure 400 {object} core.Error
// @Failure 404 {object} core.Error
// @Security JWTAuth
// @Router /v1/{route}/update [post]
func (h *handler) Update(c *fiber.Ctx) error {
	b := *requestx.Get[Update{model}](c)

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
// @Tags {tags}
// @Summary Delete a {tags} by id
// @Accept json
// @Produce json
// @Param query body Delete{model} true "query"
// @Success 200 {object} core.Success
// @Failure 400 {object} core.Error
// @Failure 404 {object} core.Error
// @Security JWTAuth
// @Router /v1/{route}/delete [post]
func (h *handler) Delete(c *fiber.Ctx) error {
	b := *requestx.Get[Delete{model}](c)

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
	tmpl = strings.ReplaceAll(tmpl, "{name}", tocase.ToLower(model))
	tmpl = strings.ReplaceAll(tmpl, "{tags}", strings.ToLower(model))
	tmpl = strings.ReplaceAll(tmpl, "{route}", strcase.ToKebab(model))

	return tmpl
}

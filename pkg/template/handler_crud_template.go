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

func (h *handler) FindById(c *fiber.Ctx) error {
	q := QueryOne{}
	_ = c.BodyParser(&q)

	if r := h.Uc.FindById(q.Id); r.Id {findOneCondition} {
		return h.Response.With(c).Response(goerror.NewOK(r))
	}

	return h.Response.With(c).Response(goerror.NewNotFound())
}

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

func (h *handler) Create(c *fiber.Ctx) error {
	b := Create{model}{}
	_ = c.BodyParser(&b)

	d, err := h.Uc.Create(&b)
	if err != nil {
		return h.Response.With(c).Response(err)
	}

	return h.Response.With(c).Response(goerror.NewCreated(d))
}

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

	return tmpl
}

package template

import (
	"fmt"
	"github.com/ettle/strcase"
)

func ValidateCrud(module string, name string) string {
	structName := strcase.ToPascal(name)
	return fmt.Sprintf(`package %s

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/fibererror"
	"github.com/prongbang/goerror"
	"`+module+`/internal/pkg/response"
)

type Validate interface {
	FindById(c *fiber.Ctx) error
	FindList(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type validate struct {
	Validate *validator.Validate
	Response fibererror.Response
}

func (v *validate) FindById(c *fiber.Ctx) error {
	body := QueryOne{}

	if err := c.BodyParser(&body); err != nil {
		return v.Response.With(c).Response(goerror.NewBadRequest())
	}
	if err := v.Validate.Struct(body); err != nil {
		return v.Response.With(c).Response(response.NewDataInvalidError())
	}

	return c.Next()
}

func (v *validate) FindList(c *fiber.Ctx) error {
	body := QueryMany{}

	if err := c.BodyParser(&body); err != nil {
		return v.Response.With(c).Response(goerror.NewBadRequest())
	}
	if err := v.Validate.Struct(body); err != nil {
		return v.Response.With(c).Response(response.NewDataInvalidError())
	}
	
	return c.Next()
}

func (v *validate) Create(c *fiber.Ctx) error {
	body := Create`+structName+`{}

	if err := c.BodyParser(&body); err != nil {
		return v.Response.With(c).Response(goerror.NewBadRequest())
	}
	if err := v.Validate.Struct(body); err != nil {
		return v.Response.With(c).Response(response.NewDataInvalidError())
	}

	return c.Next()
}

func (v *validate) Update(c *fiber.Ctx) error {
	body := Update`+structName+`{}

	if err := c.BodyParser(&body); err != nil {
		return v.Response.With(c).Response(goerror.NewBadRequest())
	}
	if err := v.Validate.Struct(body); err != nil {
		return v.Response.With(c).Response(response.NewDataInvalidError())
	}

	return c.Next()
}

func (v *validate) Delete(c *fiber.Ctx) error {
	body := Delete`+structName+`{}

	if err := c.BodyParser(&body); err != nil {
		return v.Response.With(c).Response(goerror.NewBadRequest())
	}
	if err := v.Validate.Struct(body); err != nil {
		return v.Response.With(c).Response(response.NewDataInvalidError())
	}

	return c.Next()
}	

func NewValidate(v *validator.Validate, response fibererror.Response) Validate {
	return &validate{
		Validate: v,
		Response: response,
	}
}
	`, name)
}

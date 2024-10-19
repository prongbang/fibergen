package template

import "fmt"

func Validate(name string) string {
	return fmt.Sprintf(`package %s

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/fibererror"
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
	return c.Next()
}

func (v *validate) FindList(c *fiber.Ctx) error {
	return c.Next()
}

func (v *validate) Create(c *fiber.Ctx) error {
	return c.Next()
}

func (v *validate) Update(c *fiber.Ctx) error {
	return c.Next()
}

func (v *validate) Delete(c *fiber.Ctx) error {
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

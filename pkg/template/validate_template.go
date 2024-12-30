package template

import (
	"fmt"
	"github.com/prongbang/fibergen/pkg/tocase"
)

func Validate(name string) string {
	return fmt.Sprintf(`package %s

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/fibererror"
)

type Validate interface {
	Example(c *fiber.Ctx) error
}

type validate struct {
	Validate *validator.Validate
	Response fibererror.Response
}

func (v *validate) Example(c *fiber.Ctx) error {
	return c.Next()
}

func NewValidate(v *validator.Validate, response fibererror.Response) Validate {
	return &validate{
		Validate: v,
		Response: response,
	}
}
	`, tocase.ToLower(name))
}

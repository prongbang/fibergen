package template

import "fmt"

func ValidateCrud(name string) string {
	return fmt.Sprintf(`package %s

import (
	"github.com/gofiber/fiber/v2"
)

type Validate interface {
	FindById(c *fiber.Ctx) error
	FindList(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type validate struct {
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

func NewValidate() Validate {
	return &validate{}
}
	`, name)
}

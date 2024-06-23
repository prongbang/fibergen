package template

type coreHandlerTemplate struct {
}

func (c *coreHandlerTemplate) Text() []byte {
	return []byte(`package core

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	FindById(c *fiber.Ctx) error
	FindList(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}`)
}

func CoreHandlerTemplate() Template {
	return &coreHandlerTemplate{}
}

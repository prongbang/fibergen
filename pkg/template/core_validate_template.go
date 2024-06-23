package template

type coreValidateTemplate struct {
}

func (c *coreValidateTemplate) Text() []byte {
	return []byte(`package core

import "github.com/gofiber/fiber/v2"

type Validate interface {
	FindList(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
`)
}

func CoreValidateTemplate() Template {
	return &coreValidateTemplate{}
}

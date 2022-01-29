package template

type coreHeaderTemplate struct {
}

func (c *coreHeaderTemplate) Text() []byte {
	return []byte(`package core

import (
	"github.com/gofiber/fiber/v2"
)

func AcceptLanguage(c *fiber.Ctx) string {
	return c.AcceptsLanguages()
}

func Authorization(c *fiber.Ctx) string {
	return c.Get(fiber.HeaderAuthorization)
}
`)
}

func CoreHeaderTemplate() Template {
	return &coreHeaderTemplate{}
}

package template

import (
	"strings"
)

func RequestXRequest(module string) string {
	tmpl := `package requestx

import (
	"github.com/gofiber/fiber/v2"
	"{module}/pkg/structx"
)

type model[T any] struct {
	Type T
}

func Next[T any](c *fiber.Ctx, value T) error {
	Set[T](c, value)
	return c.Next()
}

func Set[T any](c *fiber.Ctx, value T) {
	c.Locals(structx.Name(value), value)
}

func Get[T any](c *fiber.Ctx) *T {
	m := model[T]{}
	value, ok := c.Locals(structx.Name(m.Type)).(T)
	if ok {
		return &value
	}
	return nil
}`

	tmpl = strings.ReplaceAll(tmpl, "{module}", module)

	return tmpl
}

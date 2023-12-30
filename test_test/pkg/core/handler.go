package core

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	GetById(c *fiber.Ctx) error
	GetList(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
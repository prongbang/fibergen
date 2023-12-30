package core

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

var (
	StatusOK           = http.StatusText(http.StatusOK)
	StatusCreated      = http.StatusText(http.StatusCreated)
	StatusBadRequest   = http.StatusText(http.StatusBadRequest)
	StatusUnauthorized = http.StatusText(http.StatusUnauthorized)
	StatusForbidden    = http.StatusText(http.StatusForbidden)
	StatusNotFound     = http.StatusText(http.StatusNotFound)
)

type Error struct {
	Message string `json:"message"`
}

type Success struct {
	Message string `json:"message"`
}

type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Message interface{} `json:"message"`
}

type Image struct {
	URL string `json:"url"`
}

type Result struct {
	Code    int
	Message interface{}
	Data    interface{}
}

func Resp(c *fiber.Ctx, result Result) error {
	message := getMessageByStatus(result)
	switch result.Code {
	case http.StatusOK:
		return Ok(c, result.Data)
	case http.StatusCreated:
		return Created(c, result.Data)
	case http.StatusNoContent:
		return NoContent(c, message)
	case http.StatusBadRequest:
		return BadRequest(c, message)
	case http.StatusUnauthorized:
		return Unauthorized(c, message)
	case http.StatusForbidden:
		return Forbidden(c, message)
	}
	return NotFound(c, message)
}

func getMessageByStatus(result Result) interface{} {
	if result.Message == nil {
		return http.StatusText(result.Code)
	}

	// Message type
	if _, okValidate := result.Message.(validator.ValidationErrors); okValidate {
		return result.Message
	}

	msgErr, okErr := result.Message.(error)
	if okErr {
		return msgErr.Error()
	}
	return result.Message
}

func Ok(c *fiber.Ctx, data interface{}) error {
	return c.Status(http.StatusOK).JSON(data)
}

func SendStream(c *fiber.Ctx, data interface{}) error {
	return c.Status(http.StatusOK).JSON(data)
}

func Created(c *fiber.Ctx, data interface{}) error {
	return c.Status(http.StatusCreated).JSON(data)
}

func BadRequest(c *fiber.Ctx, data interface{}) error {
	if data == nil {
		return c.Status(http.StatusBadRequest).JSON(&Response{
			Message: StatusBadRequest,
		})
	}

	// Validation Errors
	if _, okValidation := data.(validator.ValidationErrors); okValidation {
		msgMap := fiber.Map{}
		for _, e := range data.(validator.ValidationErrors) {
			msgMap[e.Field()] = fiber.Map{
				"required": fmt.Sprintf("%s is %s and not empty", e.Field(), e.Tag()),
			}
		}
		return c.Status(http.StatusBadRequest).JSON(&Response{
			Message: msgMap,
		})
	}

	// Other error
	return c.Status(http.StatusBadRequest).JSON(&Response{
		Message: data,
	})
}

func NotFound(c *fiber.Ctx, data interface{}) error {
	return c.Status(http.StatusNotFound).JSON(&Response{
		Message: data,
	})
}

func NoContent(c *fiber.Ctx, data interface{}) error {
	return c.Status(http.StatusNoContent).JSON(&Response{
		Message: data,
	})
}

func Unauthorized(c *fiber.Ctx, data interface{}) error {
	return c.Status(http.StatusUnauthorized).JSON(&Response{
		Message: data,
	})
}

func Forbidden(c *fiber.Ctx, data interface{}) error {
	return c.Status(http.StatusForbidden).JSON(&Response{
		Message: data,
	})
}

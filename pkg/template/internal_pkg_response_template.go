package template

type responseTemplate struct {
}

func (c *responseTemplate) Text() []byte {
	return []byte(`package response

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/fibererror"
	"github.com/prongbang/goerror"
)

type CustomError struct {
	goerror.Body
}

// Error implements error.
func (c *CustomError) Error() string {
	return c.Message
}

func NewCustomError() error {
	return &CustomError{
		Body: goerror.Body{
			Code: "CUS001",
		},
	}
}

type customResponse struct {
}

// Response implements response.Custom.
func (c *customResponse) Response(ctx *fiber.Ctx, err error) error {
	switch resp := err.(type) {
	case *CustomError:
		return ctx.Status(http.StatusBadRequest).JSON(resp)
	}
	return nil
}

func NewCustomResponse() fibererror.Custom {
	return &customResponse{}
}

func New() fibererror.Response {
	customResp := NewCustomResponse()
	return fibererror.New(&fibererror.Config{
		Custom: &customResp,
	})
}`)
}

func InternalPkgResponseTemplate() Template {
	return &responseTemplate{}
}

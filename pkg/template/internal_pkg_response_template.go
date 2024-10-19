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

type CommitError struct {
	goerror.Body
}

// Error implements error.
func (c *CommitError) Error() string {
	return c.Message
}

func NewCommitError() error {
	return &CommitError{
		Body: goerror.Body{
			Code:    "DTB001",
			Message: "Transaction commit failed",
		},
	}
}

type InsertError struct {
	goerror.Body
}

// Error implements error.
func (c *InsertError) Error() string {
	return c.Message
}

func NewInsertError() error {
	return &InsertError{
		Body: goerror.Body{
			Code:    "DTB002",
			Message: "Failed to insert a child row",
		},
	}
}

type UpdateError struct {
	goerror.Body
}

// Error implements error.
func (c *UpdateError) Error() string {
	return c.Message
}

func NewUpdateError() error {
	return &UpdateError{
		Body: goerror.Body{
			Code:    "DTB003",
			Message: "Failed to update a child row",
		},
	}
}

type DeleteError struct {
	goerror.Body
}

// Error implements error.
func (c *DeleteError) Error() string {
	return c.Message
}

func NewDeleteError() error {
	return &DeleteError{
		Body: goerror.Body{
			Code:    "DTB004",
			Message: "Failed to delete a child row",
		},
	}
}

type DataInvalidError struct {
	goerror.Body
}

// Error implements error.
func (c *DataInvalidError) Error() string {
	return c.Message
}

func NewDataInvalidError() error {
	return &DataInvalidError{
		Body: goerror.Body{
			Code:    "CLE029",
			Message: "Invalid data provided",
		},
	}
}

type customResponse struct {
}

// Response implements response.Custom.
func (c *customResponse) Response(ctx *fiber.Ctx, err error) error {
	switch resp := err.(type) {
	case *UpdateError, *DeleteError, *CommitError, *InsertError, *DataInvalidError:
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

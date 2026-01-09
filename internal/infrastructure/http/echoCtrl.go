// Package http provides methods to handle requests
package http

import (
	"github.com/gofiber/fiber/v2"
)

type EchoController struct {
}

func (ctrl *EchoController) Handler(ctx *fiber.Ctx) error {
	records, err := readRequest(ctx)
	if err != nil {
		return HandleHttpError(ctx, err)
	}

	ctx.Status(fiber.StatusOK).SendString(convertToResponse(records))
	return nil
}

func NewEchoController() *EchoController {
	return &EchoController{}
}

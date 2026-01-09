// Package http provides methods to handle requests
package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ibanezv/league_challenge/internal/domain"
)

type InvertController struct {
	invertUseCase domain.InvertUseCase
}

func (ctrl *InvertController) Handler(ctx *fiber.Ctx) error {
	records, err := readRequest(ctx)
	if err != nil {
		return HandleHttpError(ctx, err)
	}

	result := ctrl.invertUseCase.Do(ctx.UserContext(), records)

	return ctx.Status(fiber.StatusOK).SendString(convertToResponse(result))
}

func NewInvertController(invertuc domain.InvertUseCase) *InvertController {
	return &InvertController{invertUseCase: invertuc}
}

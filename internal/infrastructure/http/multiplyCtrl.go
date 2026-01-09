// Package http provides methods to handle requests
package http

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ibanezv/league_challenge/internal/domain"
)

type MultiplyController struct {
	multiplyUseCase domain.MultiplyUseCase
}

func (ctrl *MultiplyController) Handler(ctx *fiber.Ctx) error {
	input, err := readRequest(ctx)
	if err != nil {
		return HandleHttpError(ctx, err)
	}

	records, err := convertToInt(input)
	if err != nil {
		return HandleHttpError(ctx, err)
	}

	resp := ctrl.multiplyUseCase.Do(ctx.UserContext(), records)
	return ctx.Status(fiber.StatusOK).SendString(strconv.Itoa(resp))
}

func NewMultiplyController(multiplyuc domain.MultiplyUseCase) *MultiplyController {
	return &MultiplyController{multiplyUseCase: multiplyuc}
}

// Package http provides methods to handle requests
package http

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ibanezv/league_challenge/internal/domain"
)

type FlattenController struct {
	flatenUseCase domain.FlattenUseCase
}

func (ctrl *FlattenController) Handler(ctx *fiber.Ctx) error {
	records, err := readRequest(ctx)
	if err != nil {
		return HandleHttpError(ctx, err)
	}

	result := ctrl.flatenUseCase.Do(ctx.UserContext(), records)
	return ctx.Status(fiber.StatusOK).SendString(strings.Join(result, ","))
}

func NewFlattenController(flatenUc domain.FlattenUseCase) *FlattenController {
	return &FlattenController{flatenUseCase: flatenUc}
}

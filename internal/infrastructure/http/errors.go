package http

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrBadFileInput        = errors.New("File provides is not valid")
	ErrBadContentFileInput = errors.New("Matrix provides is not valid")
	ErrMatrixEmpty         = errors.New("Matrix provides is empty")
	ErrMatrixNoSquare      = errors.New("Matrix provided is not a square")
)

func HandleHttpError(ctx *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, ErrBadFileInput):
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	case errors.Is(err, ErrBadContentFileInput):
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	default:
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
}

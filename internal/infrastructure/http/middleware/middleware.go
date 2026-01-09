// Package middleware provides methods for input validation
package middleware

import (
	"encoding/csv"
	"path/filepath"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const (
	headerName = "file"
	separator  = ','
	extension  = ".csv"
)

// ValidationInput validate if request contains a matrix of integers
func ValidationInput(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile(headerName)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON("File provides is not valid")
	}

	uploadedFile, err := file.Open()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON("File provides is not valid")
	}
	defer uploadedFile.Close()

	reader := csv.NewReader(uploadedFile)
	reader.Comma = separator
	records, err := reader.ReadAll()

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON("File provides is not valid")
	}

	for _, row := range records {
		for _, col := range row {
			if _, err := strconv.Atoi(col); err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON("Matrix provides is not valid")
			}
		}
	}
	return ctx.Next()
}

// ValidationMatrix validate if request contains a valid file
func ValidationMatrix(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile(headerName)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON("File provides is not valid")
	}

	if filepath.Ext(file.Filename) != extension {
		return ctx.Status(fiber.StatusBadRequest).JSON("File provides is not valid")
	}

	uploadedFile, err := file.Open()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON("File provides is not valid")
	}
	defer uploadedFile.Close()

	reader := csv.NewReader(uploadedFile)
	reader.Comma = separator
	records, err := reader.ReadAll()

	if len(records) == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON("Matrix provides is empty")
	}

	if len(records) != len(records[0]) {
		return ctx.Status(fiber.StatusBadRequest).JSON("Matrix provided is not a square")
	}

	return ctx.Next()
}

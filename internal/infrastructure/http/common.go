package http

import (
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const separator = ','

func readRequest(ctx *fiber.Ctx) ([][]string, error) {
	file, err := ctx.FormFile("file")
	if err != nil {
		return nil, ErrBadFileInput
	}

	uploadedFile, err := file.Open()
	if err != nil {
		return nil, ErrBadFileInput
	}
	defer uploadedFile.Close()

	reader := csv.NewReader(uploadedFile)
	reader.Comma = separator
	records, err := reader.ReadAll()
	if err != nil {
		return nil, ErrBadFileInput
	}

	return records, nil
}

func convertToInt(matrix [][]string) ([][]int, error) {
	resp := make([][]int, 0)
	for i := 0; i <= len(matrix)-1; i++ {
		row := make([]int, 0)
		for j := 0; j <= len(matrix[i])-1; j++ {
			element, err := strconv.Atoi(matrix[i][j])
			if err != nil {
				return nil, ErrBadContentFileInput
			}
			row = append(row, element)
		}
		resp = append(resp, row)
	}
	return resp, nil
}

func convertToResponse(matrix [][]string) string {
	var response string
	for _, row := range matrix {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	return response
}

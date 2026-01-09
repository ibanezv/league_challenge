package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Given: different scenarios.
// When: do sum use case.
// Then: retrieves expected result.
func TestSumUseCase(t *testing.T) {
	var cases = []struct {
		name     string
		input    [][]int
		expected int
	}{
		{
			name:     "sum success",
			input:    [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expected: 45,
		},
	}
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			useCase := NewSumUseCase()

			// Act
			resp := useCase.Do(context.TODO(), test.input)

			// Assert
			assert.Equal(t, test.expected, resp)
		})

	}
}

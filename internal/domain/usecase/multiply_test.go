package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Given: different scenarios.
// When: do multiply use case.
// Then: retrieves expected result.
func TestMultiplyUseCase(t *testing.T) {
	var cases = []struct {
		name     string
		input    [][]int
		expected int
	}{
		{
			name:     "multiply success",
			input:    [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expected: 362880,
		},
	}
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			useCase := NewMultiplyUseCase()

			// Act
			resp := useCase.Do(context.TODO(), test.input)

			// Assert
			assert.Equal(t, test.expected, resp)
		})

	}
}

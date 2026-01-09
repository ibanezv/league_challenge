package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Given: different scenarios.
// When: do flatten use case.
// Then: retrieves expected result.
func TestFlattenUseCase(t *testing.T) {
	var cases = []struct {
		name     string
		input    [][]string
		expected []string
	}{
		{
			name:     "flatten success",
			input:    [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
			expected: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
		},
	}
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			useCase := NewFlattenUseCase()

			// Act
			resp := useCase.Do(context.TODO(), test.input)

			// Assert
			assert.Equal(t, test.expected, resp)
		})

	}
}

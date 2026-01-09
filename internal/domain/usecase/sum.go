// Package usecase provides logic defined for code challeng
package usecase

import (
	"context"

	"github.com/ibanezv/league_challenge/internal/domain"
)

type sum struct {
}

func (s *sum) Do(ctx context.Context, arr [][]int) int {
	result := 0
	for i := 0; i <= len(arr)-1; i++ {
		for j := 0; j <= len(arr[i])-1; j++ {
			result += arr[i][j]
		}
	}
	return result
}

func NewSumUseCase() domain.SumUseCase {
	return &sum{}
}

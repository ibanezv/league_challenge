// Package usecase provides logic defined for code challeng
package usecase

import (
	"context"

	"github.com/league/league_challenge/internal/domain"
)

type invert struct {
}

func (i *invert) Do(ctx context.Context, arr [][]string) [][]string {
	resp := make([][]string, 0)
	for i := 0; i <= len(arr)-1; i++ {
		row := make([]string, 0)
		for j := 0; j <= len(arr[i])-1; j++ {
			row = append(row, arr[j][i])
		}
		resp = append(resp, row)
	}
	return resp
}

func NewInvertUseCase() domain.InvertUseCase {
	return &invert{}
}

// Package usecase provides logic defined for code challeng
package usecase

import (
	"context"

	"github.com/ibanezv/league_challenge/internal/domain"
)

type flatten struct {
}

func (f *flatten) Do(ctx context.Context, arr [][]string) []string {
	resp := make([]string, 0)
	for i := 0; i <= len(arr)-1; i++ {
		for j := 0; j <= len(arr[i])-1; j++ {
			resp = append(resp, arr[i][j])
		}
	}
	return resp
}

func NewFlattenUseCase() domain.FlattenUseCase {
	return &flatten{}
}

package domain

import "context"

type FlattenUseCase interface {
	Do(context.Context, [][]string) []string
}

type InvertUseCase interface {
	Do(context.Context, [][]string) [][]string
}

type MultiplyUseCase interface {
	Do(context.Context, [][]int) int
}
type SumUseCase interface {
	Do(context.Context, [][]int) int
}

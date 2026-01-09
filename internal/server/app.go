// Package infrastructure provides features to run server
package server

import (
	"context"
	"log"

	"github.com/ibanezv/league_challenge/internal/domain/usecase"
	"github.com/ibanezv/league_challenge/internal/infrastructure/http"
	"github.com/ibanezv/league_challenge/internal/server/infrastructure"
	"go.uber.org/fx"
)

func initialize(
	app *infrastructure.AppServer,
	ctx context.Context,
) {
	if err := app.Listen(); err != nil {
		log.Fatal("failure at web server startup")
	}
}

func InitServer() {
	fx.New(
		NewServer(),
		fx.Invoke(initialize),
	).Run()
}

func NewServer() fx.Option {
	injections := []interface{}{
		infrastructure.NewAppServer,
	}

	injections = append(injections, getDomainInjector()...)
	injections = append(injections, getInfraInjector()...)
	return fx.Provide(injections...)
}

func getDomainInjector() []interface{} {
	return []interface{}{
		usecase.NewFlattenUseCase,
		usecase.NewInvertUseCase,
		usecase.NewMultiplyUseCase,
		usecase.NewSumUseCase,
	}
}

func getInfraInjector() []interface{} {
	return []interface{}{
		context.Background,
		http.NewEchoController,
		http.NewFlattenController,
		http.NewInvertController,
		http.NewMultiplyController,
		http.NewSumController,
	}
}

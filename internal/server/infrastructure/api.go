// Package infrastructure provides features to expose api
package infrastructure

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/ibanezv/league_challenge/internal/infrastructure/http"
	"github.com/ibanezv/league_challenge/internal/infrastructure/http/middleware"
)

type AppServer struct {
	echoController     *http.EchoController
	flattenController  *http.FlattenController
	invertController   *http.InvertController
	multiplyController *http.MultiplyController
	sumController      *http.SumController
}

func (s *AppServer) Listen() error {
	f := fiber.New()

	f.Use(recover.New())

	f.Post("/echo", middleware.ValidationMatrix, middleware.ValidationInput, s.echoController.Handler)
	f.Post("/flatten", middleware.ValidationMatrix, s.flattenController.Handler)
	f.Post("/invert", middleware.ValidationMatrix, s.invertController.Handler)
	f.Post("/multiply", middleware.ValidationMatrix, s.multiplyController.Handler)
	f.Post("/sum", middleware.ValidationMatrix, s.sumController.Handler)

	return f.Listen(":8080")
}

func NewAppServer(
	echoController *http.EchoController,
	flattenController *http.FlattenController,
	invertController *http.InvertController,
	multiplyController *http.MultiplyController,
	sumController *http.SumController,
) *AppServer {
	return &AppServer{echoController: echoController,
		flattenController:  flattenController,
		invertController:   invertController,
		multiplyController: multiplyController,
		sumController:      sumController}
}

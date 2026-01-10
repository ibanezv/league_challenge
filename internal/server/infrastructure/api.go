// Package infrastructure provides features to expose api
package infrastructure

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/league/league_challenge/internal/infrastructure/http"
	"github.com/league/league_challenge/internal/infrastructure/http/middleware"
)

type AppServer struct {
	echoController     *http.EchoController
	flattenController  *http.FlattenController
	invertController   *http.InvertController
	multiplyController *http.MultiplyController
	sumController      *http.SumController
}

func (s *AppServer) Listen() error {
	api := fiber.New()

	api.Use(recover.New())

	group := api.Group("/api/v1")
	// swagger
	group.Get("/swagger/*", swagger.HandlerDefault)

	group.Post("/echo", middleware.ValidationMatrix, s.echoController.Handler)
	group.Post("/flatten", middleware.ValidationMatrix, s.flattenController.Handler)
	group.Post("/invert", middleware.ValidationMatrix, s.invertController.Handler)
	group.Post("/multiply", middleware.ValidationMatrix, middleware.ValidationInput, s.multiplyController.Handler)
	group.Post("/sum", middleware.ValidationMatrix, middleware.ValidationInput, s.sumController.Handler)

	return api.Listen(":8080")
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

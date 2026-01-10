package main

import (
	_ "github.com/league/league_challenge/cmd/docs"

	"github.com/league/league_challenge/internal/server"
)

// @title League Code Challenge
// @version 1.0
// @description Leagge Code Challenge.
// @host localhost:8080
// @BasePath /api/v1
// @schemes http
func main() {
	server.InitServer()
}

package main

import (
	"github.com/eddyvy/template/internal/player"
	"github.com/eddyvy/template/internal/referee"
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "UP"})
	})

	app.Get("/players", player.GetPlayers)
	app.Get("/players/:id", player.GetPlayer)
	app.Post("/players", player.CreatePlayer)
	app.Put("/players/:id", player.UpdatePlayer)
	app.Delete("/players/:id", player.DeletePlayer)

	app.Get("/referees", referee.GetReferees)
	app.Get("/referees/:id", referee.GetReferee)
	app.Post("/referees", referee.CreateReferee)
	app.Put("/referees/:id", referee.UpdateReferee)
	app.Delete("/referees/:id", referee.DeleteReferee)
}

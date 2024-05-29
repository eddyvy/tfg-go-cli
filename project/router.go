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

	app.Get("/players", player.HandleGetAll)
	app.Get("/players/:id", player.HandleGetOne)
	app.Post("/players", player.HandlePost)
	app.Put("/players/:id", player.HandlePut)
	app.Delete("/players/:id", player.HandleDelete)

	app.Get("/referees", referee.HandleGetAll)
	app.Get("/referees/:id", referee.HandleGetOne)
	app.Post("/referees", referee.HandlePost)
	app.Put("/referees/:id", referee.HandlePut)
	app.Delete("/referees/:id", referee.HandleDelete)
}

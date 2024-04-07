package main

import (
	"log"

	"github.com/eddyvy/template/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	// Initialize the database
	database.InitDb()
}

func main() {
	// Create a new Fiber app
	app := fiber.New()

	// Apply middlewares
	app.Use(logger.New())

	// Apply routes
	SetRoutes(app)

	// Start server
	log.Fatal(app.Listen("127.0.0.1:3000"))
}

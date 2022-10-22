package main

import (
	"github.com/NeoTRAN001/hello-fiber/AnimeHandler"
	"github.com/NeoTRAN001/hello-fiber/UserHandler"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes(app)

	app.Listen(":3000")
}

func routes(app *fiber.App) {
	UserHandler.Routes(app)
	AnimeHandler.Routes(app)
}

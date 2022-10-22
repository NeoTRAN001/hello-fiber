package AnimeHandler

import (
	"github.com/NeoTRAN001/hello-fiber/AnimeHandler/controllers"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/user", controllers.GetAllAnime)
}

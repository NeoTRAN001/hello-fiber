package UserHandler

import (
	"github.com/NeoTRAN001/hello-fiber/UserHandler/controllers"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/user", controllers.HandleUser)
	app.Post("users", controllers.HandleCreateUser)
}

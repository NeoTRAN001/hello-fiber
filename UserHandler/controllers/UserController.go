package controllers

import (
	"github.com/NeoTRAN001/hello-fiber/UserHandler/dto"
	"github.com/NeoTRAN001/hello-fiber/UserHandler/services"
	"github.com/gofiber/fiber/v2"
)

func HandleUser(c *fiber.Ctx) error {

	var user = dto.UserDTO{
		Firstname: "Joe",
		Lastname:  "Doe",
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func HandleCreateUser(c *fiber.Ctx) error {
	user := dto.UserDTO{}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseDTO{
			Status:  fiber.StatusBadRequest,
			Message: "Format request incorrect",
		})
	}

	err, response := services.CreateUser(user)

	if err != nil {
		return c.Status(response.Status).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

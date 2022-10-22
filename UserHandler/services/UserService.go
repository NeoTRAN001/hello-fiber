package services

import (
	"github.com/NeoTRAN001/hello-fiber/UserHandler/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateUser(user dto.UserDTO) (error, dto.ResponseDTO) {

	if user.Firstname == "" || user.Lastname == "" {
		return fiber.ErrBadRequest, dto.ResponseDTO{
			Status:  fiber.StatusBadRequest,
			Message: "Required firstname and lastname",
		}
	}

	var id string = uuid.NewString()

	return nil, dto.ResponseDTO{
		Status:  fiber.StatusAccepted,
		Message: "Create User with ID: " + id,
	}
}

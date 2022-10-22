package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func GetAllAnime(c *fiber.Ctx) error {
	return c.SendString("Hello from Anime Controller!")
}

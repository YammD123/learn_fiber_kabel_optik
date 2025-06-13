package controllers

import (
	"myapp/services"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	return c.JSON(services.GetAllUsers())
}
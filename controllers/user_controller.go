package controllers

import (
	"myapp/models"
	"myapp/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllUsers(c *fiber.Ctx) error {
	return c.JSON(services.GetAllUsers())
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	user.ID = uuid.New().String()
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"+err.Error()})
	}
	cratedUser, err := services.CreateUser(&user)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request" + err.Error()})
	}
	return c.JSON(cratedUser)
}
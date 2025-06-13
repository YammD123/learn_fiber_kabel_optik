package routes

import (
	"myapp/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	user:=app.Group("/user")

	user.Get("/",controllers.GetAllUsers)
}
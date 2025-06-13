package main

import (
	"myapp/routes"

	"github.com/gofiber/fiber/v2"
)

func main(){
	app:=fiber.New()
	routes.UserRoute(app)
	app.Listen(":3000")
}

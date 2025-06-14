package main

import (
	"myapp/configs"
	"myapp/routes"

	"github.com/gofiber/fiber/v2"
)

func main(){
	app:=fiber.New()
	configs.ConnectDatabase()
	routes.UserRoute(app)
	app.Listen(":3000")
}

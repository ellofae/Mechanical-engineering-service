package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ellofae/Mechanical-engineering-service/app/controllers"
)

func main() {
    app := fiber.New()

    app.Get("/services", controllers.GetServices)
	app.Post("/services", controllers.CreateService)

    app.Listen(":3000")
}
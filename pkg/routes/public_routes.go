package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ellofae/Mechanical-engineering-service/app/controllers"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/mechanics")

	route.Get("/services", controllers.GetServices)
	route.Get("/service/:id", controllers.GetService)
}
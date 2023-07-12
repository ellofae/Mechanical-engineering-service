package routes

import (
	"github.com/ellofae/Mechanical-engineering-service/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/mechanics")

	route.Get("/services", controllers.GetServices)
	route.Get("/service", controllers.GetService)
}

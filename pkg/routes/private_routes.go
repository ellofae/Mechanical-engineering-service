package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ellofae/Mechanical-engineering-service/app/controllers"
)

func PrivateRoutes(a *fiber.App) {
	route := a.Group("/mechanics")

	route.Post("/service/create", controllers.CreateService)
	route.Put("/service/update", controllers.UpdateService)
	route.Delete("/service/delete", controllers.DeleteService)
}
package routes

import (
	"github.com/ellofae/Mechanical-engineering-service/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	route := a.Group("/mechanics")

	route.Post("/service/create", controllers.CreateService)
	route.Put("/service/update", controllers.UpdateService)
	route.Delete("/service/delete", controllers.DeleteService)

	route.Post("/vehicle/create", controllers.CreateVehicle)
	route.Post("/vehicle/update", controllers.UpdateVehicle)
	route.Post("/vehicle/delete", controllers.DeleteVehicle)
}

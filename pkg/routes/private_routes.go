package routes

import (
	"github.com/ellofae/Mechanical-engineering-service/app/controllers"
	"github.com/ellofae/Mechanical-engineering-service/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	route := a.Group("/mechanics")

	route.Post("/service/create", middleware.UserIdentity(), controllers.CreateService)
	route.Put("/service/update", middleware.UserIdentity(), controllers.UpdateService)
	route.Delete("/service/delete", middleware.UserIdentity(), controllers.DeleteService)

	route.Post("/vehicle/create", middleware.UserIdentity(), controllers.CreateVehicle)
	route.Post("/vehicle/update", middleware.UserIdentity(), controllers.UpdateVehicle)
	route.Post("/vehicle/delete", middleware.UserIdentity(), controllers.DeleteVehicle)
}

package routes

import (
	"github.com/ellofae/Mechanical-engineering-service/app/controllers"
	"github.com/ellofae/Mechanical-engineering-service/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/mechanics")

	route.Get("/services", middleware.AuthenticateUser, controllers.GetServices)
	route.Get("/service", middleware.AuthenticateUser, controllers.GetService)

	route.Get("/vehicles", middleware.AuthenticateUser, controllers.GetVehicles)
	route.Get("/vehicle", middleware.AuthenticateUser, controllers.GetVehicle)
}

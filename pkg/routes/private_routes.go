package routes

import (
	"github.com/ellofae/Mechanical-engineering-service/app/controllers"
)

func PrivateRoutes(a *fiber.App) {
	route := a.Group("/mechanics")

	route.Get("/service/create", controllers.CreateService)
	route.Get("/service/update", controllers.UpdateService)
}
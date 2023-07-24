package routes

import (
	"github.com/ellofae/Mechanical-engineering-service/app/controllers"
	"github.com/ellofae/Mechanical-engineering-service/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	route := a.Group("/mechanics")

	// route.Post("/service/create", middleware.UserIdentity(), controllers.CreateService)   // middleware is used
	// route.Put("/service/update", middleware.UserIdentity(), controllers.UpdateService)    // middleware is used
	// route.Delete("/service/delete", middleware.UserIdentity(), controllers.DeleteService) // middleware is used

	route.Post("/service/create", middleware.AuthenticateUser, controllers.CreateService)   // middleware is used
	route.Put("/service/update", middleware.AuthenticateUser, controllers.UpdateService)    // middleware is used
	route.Delete("/service/delete", middleware.AuthenticateUser, controllers.DeleteService) // middleware is used

	// route.Post("/vehicle/create", controllers.CreateVehicle) // cookies are used
	// route.Post("/vehicle/update", controllers.UpdateVehicle) // cookies are used
	// route.Post("/vehicle/delete", controllers.DeleteVehicle) // cookies are used

	route.Post("/vehicle/create", middleware.AuthenticateUser, controllers.CreateVehicle) // cookies are used
	route.Post("/vehicle/update", middleware.AuthenticateUser, controllers.UpdateVehicle) // cookies are used
	route.Post("/vehicle/delete", middleware.AuthenticateUser, controllers.DeleteVehicle) // cookies are used
}

package routes

import (
	"github.com/ellofae/Mechanical-engineering-service/app/controllers"
	"github.com/ellofae/Mechanical-engineering-service/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func FileloaderRoutes(a *fiber.App) {
	route := a.Group("/fileloader")

	route.Get("/load", middleware.AuthenticateUser, controllers.LoadingFileFields)
	route.Post("/load", middleware.AuthenticateUser, controllers.LoadFile)

	route.Get("/getfile", middleware.AuthenticateUser, controllers.AccessingFile)
	route.Post("/getfile", middleware.AuthenticateUser, controllers.AccessFile)
}

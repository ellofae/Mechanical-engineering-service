package routes

import (
	"github.com/ellofae/Mechanical-engineering-service/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func FileloaderRoutes(a *fiber.App) {
	route := a.Group("/fileloader")

	route.Get("/load", controllers.LoadingFileFields)
	route.Post("/load", controllers.LoadFile)

	route.Get("/getfile", controllers.AccessingFile)
	route.Post("/getfile", controllers.AccessFile)
}

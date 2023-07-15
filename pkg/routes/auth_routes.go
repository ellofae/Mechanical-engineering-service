package routes

import (
	"github.com/ellofae/Mechanical-engineering-service/pkg/auth/controllers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(a *fiber.App) {
	route := a.Group("/auth")

	route.Get("/signin", controllers.SignIn)
	route.Get("/signup", controllers.SingUp)
}

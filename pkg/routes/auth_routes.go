package routes

import (
	"github.com/ellofae/Mechanical-engineering-service/pkg/auth/controllers"
	"github.com/ellofae/Mechanical-engineering-service/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(a *fiber.App) {
	route := a.Group("/auth")

	route.Post("/signin", controllers.SignIn)
	route.Post("/signup", controllers.SingUp)

	route.Get("/user", middleware.UserIdentity(), controllers.GetUser)
	route.Get("/users", middleware.UserIdentity(), controllers.GetUsers)
}

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
	route.Post("/logout", controllers.Logout)

	route.Post("/refresh", middleware.AuthenticateUser, controllers.RefreshToken)

	route.Get("/user", middleware.AuthenticateUser, controllers.GetUser)
	route.Get("/users", middleware.AuthenticateUser, controllers.GetUsers)
}

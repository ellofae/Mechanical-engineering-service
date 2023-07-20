package routes

import (
	"github.com/ellofae/Mechanical-engineering-service/pkg/auth/controllers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(a *fiber.App) {
	route := a.Group("/auth")

	route.Get("/signup", controllers.RegisterUser)
	route.Get("/signin", controllers.LoginrUser)

	route.Post("/signin", controllers.SignIn)
	route.Post("/signup", controllers.SingUp)
	route.Post("/logout", controllers.Logout)

	route.Get("/user", controllers.GetUser)   // cookies are used
	route.Get("/users", controllers.GetUsers) // cookies are used
}

package routes

import (
	"github.com/gofiber/fiber/v2"
	swagger "github.com/arsmn/fiber-swagger/v2"
)

func SwaggerRoute(a *fiber.App) {
	route := a.Group("/docs")

	route.Get("*", swagger.HandlerDefault)
}
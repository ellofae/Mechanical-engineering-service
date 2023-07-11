package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/ellofae/Mechanical-engineering-service/app/controllers"
)

func SwaggerRoute(a *fiber.App) {
	route := a.Group("/docs")

	route.Get("*", swagger.HandlerDefault)
}
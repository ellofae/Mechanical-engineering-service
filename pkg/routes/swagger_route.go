package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "github.com/ellofae/Mechanical-engineering-service/docs"
)

func SwaggerRoute(a *fiber.App) {
	a.Get("/swagger/*", swagger.HandlerDefault)
}

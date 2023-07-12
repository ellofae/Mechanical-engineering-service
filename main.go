package main

import (
	"github.com/ellofae/Mechanical-engineering-service/pkg/configs"
	"github.com/ellofae/Mechanical-engineering-service/pkg/middleware"
	"github.com/ellofae/Mechanical-engineering-service/pkg/routes"
	"github.com/ellofae/Mechanical-engineering-service/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(configs.NewConfig())

	middleware.FiberMiddleware(app)

	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)

	utils.StartServerWithGracefulShutdown(app)
}

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ellofae/Mechanical-engineering-service/pkg/utils"
	"github.com/ellofae/Mechanical-engineering-service/pkg/configs"
	"github.com/ellofae/Mechanical-engineering-service/pkg/routes"
)

func main() {
    app := fiber.New(configs.NewConfig())

	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	routes.SwaggerRoute(app)

    utils.StartServerWithGracefulShutdown(app)
}
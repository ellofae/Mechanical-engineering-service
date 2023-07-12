package main

import (
	"github.com/ellofae/Mechanical-engineering-service/pkg/configs"
	"github.com/ellofae/Mechanical-engineering-service/pkg/middleware"
	"github.com/ellofae/Mechanical-engineering-service/pkg/routes"
	"github.com/ellofae/Mechanical-engineering-service/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:5000
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	app := fiber.New(configs.NewConfig())

	middleware.FiberMiddleware(app)

	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	routes.SwaggerRoute(app)

	utils.StartServerWithGracefulShutdown(app)
}

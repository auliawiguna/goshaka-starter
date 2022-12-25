package main

import (
	appMiddleware "goshaka/app/middlewares"
	apiRoutes "goshaka/app/routes"
	appConfig "goshaka/configs"
	appDatabase "goshaka/database"
	appHelper "goshaka/helpers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//Database
	appDatabase.Connect()

	//Apply default middleware
	appMiddleware.DefaultMiddleware(app)

	//Router
	apiRoutes.MainRoutes(app)
	apiRoutes.ApiRoutes(app)

	if appConfig.GetEnv("ENV") == "dev" {
		appHelper.StartServer(app)
	} else {
		appHelper.StartServerWithGracefulShutdown(app)
	}
}

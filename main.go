package main

import (
	appMiddleware "goshaka/app/middlewares"
	apiRoutes "goshaka/app/routes"
	appConfig "goshaka/configs"
	appDatabase "goshaka/database"
	appHelper "goshaka/helpers"

	"github.com/gofiber/fiber/v2"
)

// @title Goshaka Golang API Starter
// @version 1.0
// @Description This is a API boilerplate using Golang
// @contact.name Aulia Wiguna
// @contact.url https://github.com/auliawiguna/
// @contact.email wigunaahmadaulia@gmail.com
// @host 127.0.0.1:3000
// @BasePath /
// @schemas http
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

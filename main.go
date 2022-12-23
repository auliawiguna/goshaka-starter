package main

import (
	apiRoutes "goshaka/app/routes"
	appConfig "goshaka/configs"
	appDatabase "goshaka/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	port := appConfig.GetEnv("PORT")

	//Database
	appDatabase.Connect()

	//Router
	apiRoutes.MainRoutes(app)
	apiRoutes.ApiRoutes(app)

	app.Listen(":" + port)
}

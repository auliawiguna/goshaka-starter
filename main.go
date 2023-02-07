package main

import (
	appMiddleware "goshaka/app/middlewares"
	apiRoutes "goshaka/app/routes"
	appConfig "goshaka/configs"
	appDatabase "goshaka/database"
	appHelper "goshaka/helpers"
	"goshaka/jobs"

	"github.com/gofiber/fiber/v2"
)

// @title Goshaka Golang API Starter
// @version 1.0
// @Description This is an API boilerplate using Golang
// @contact.name Aulia Wiguna
// @contact.url https://github.com/auliawiguna/
// @contact.email wigunaahmadaulia@gmail.com
// @host 127.0.0.1:3000
// @BasePath /
// @schemas http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	config := appConfig.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Database
	errConn := appDatabase.Connect()
	if errConn != nil {
		panic("cannot connect database")
	}

	db, err := appDatabase.DB.DB()
	if err != nil {
		db.Close()
		panic("cannot connect database")
	}
	defer db.Close()

	// Apply default middleware
	appMiddleware.DefaultMiddleware(app)

	// Router
	apiRoutes.MainRoutes(app)
	apiRoutes.ApiRoutes(app)
	apiRoutes.StaticFile(app)

	// Run cronjob
	jobs.RunCron()

	if appConfig.GetEnv("ENV") == "dev" {
		appHelper.StartServer(app)
	} else {
		appHelper.StartServerWithGracefulShutdown(app)
	}
}

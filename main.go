package main

import (
	"log"
	"os"

	apiRoutes "goshaka/app/routes"
	appConfig "goshaka/configs"
	appDatabase "goshaka/database"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func getEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error Loading Env")
	}

	return os.Getenv(key)
}

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

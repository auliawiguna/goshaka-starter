package routes

import (
	routeV1 "goshaka/app/routes/api/v1"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func MainRoutes(app *fiber.App) {
	mainRoute := app.Group("/", logger.New())
	mainRoute.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("Hi")
		return err
	})
}

func ApiRoutes(app *fiber.App) {
	apiV1 := app.Group("/api/v1", logger.New())

	routeV1.IndexRoute(apiV1)
	routeV1.NoteRoute(apiV1)
}
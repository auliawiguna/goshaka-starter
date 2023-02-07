package routes

import (
	routeV1 "goshaka/app/routes/api/v1"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func StaticFile(app *fiber.App) {
	// Serve local path
	app.Static("/images", "./public/images", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    false,
		Index:     "index.html",
		MaxAge:    3600,
	})
}

func MainRoutes(app *fiber.App) {
	mainRoute := app.Group("/", logger.New())
	mainRoute.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("Hi")
		return err
	})
}

func ApiRoutes(app *fiber.App) {
	apiV1 := app.Group("/api/v1", logger.New())

	// Sample of protected route
	routeV1.IndexProtectedRoute(apiV1)

	routeV1.IndexRoute(apiV1)
	routeV1.NoteRoute(apiV1)
	routeV1.AuthRoute(apiV1)
	routeV1.RoleRoute(apiV1)
	routeV1.PermissionRoute(apiV1)
	routeV1.UserRoute(apiV1)
	routeV1.SwaggerRoute(apiV1)
}

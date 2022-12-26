package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func DefaultMiddleware(a *fiber.App) {
	a.Use(
		cors.New(),
		logger.New(),
		compress.New(compress.Config{
			Level: compress.LevelBestSpeed,
		}),
	)
}

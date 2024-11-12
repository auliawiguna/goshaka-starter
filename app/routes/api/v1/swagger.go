package v1

import (
	_ "goshaka/docs"

	swagger "github.com/gofiber/swagger"

	"github.com/gofiber/fiber/v2"
)

func SwaggerRoute(router fiber.Router) {
	route := router.Group("/")
	route.Get("documentation/*", swagger.HandlerDefault)
}

package v1

import (
	"github.com/gofiber/fiber/v2"
)

func IndexRoute(router fiber.Router) {
	route := router.Group("/")
	route.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("Hi There!")
		return err
	})
}

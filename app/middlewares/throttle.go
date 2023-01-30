package middlewares

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func ThrottleByKey(key string, max int, sec int) func(c *fiber.Ctx) error {
	return limiter.New(limiter.Config{
		// Next: func(c *fiber.Ctx) bool {
		// return c.IP() == "127.0.0.1"
		// },
		Max:        max,
		Expiration: time.Duration(sec) * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return key
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(http.StatusTooManyRequests).JSON(fiber.Map{
				"error":   true,
				"message": "Too Many Request",
			})
		},
	})
}

func ThrottleByIp(max int, sec int) func(c *fiber.Ctx) error {
	return limiter.New(limiter.Config{
		// Next: func(c *fiber.Ctx) bool {
		// return c.IP() == "127.0.0.1"
		// },
		Max:        max,
		Expiration: time.Duration(sec) * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(http.StatusTooManyRequests).JSON(fiber.Map{
				"error":   true,
				"message": "Too Many Request",
			})
		},
	})
}

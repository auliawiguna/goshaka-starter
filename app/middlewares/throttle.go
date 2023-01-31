package middlewares

import (
	"goshaka/helpers"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

// HTTP Throttle by key
//
//	param	key string
//	param	max int
//	param	sec int
//	receiver c *fiber.Ctx
//	return error
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
			return helpers.TooManyRequestResponse(c)
		},
	})
}

// HTTP Throttle by key and user's IP address
//
//	param	key string
//	param	max int
//	param	sec int
//	receiver c *fiber.Ctx
//	return error
func ThrottleByKeyAndIP(key string, max int, sec int) func(c *fiber.Ctx) error {
	return limiter.New(limiter.Config{
		// Next: func(c *fiber.Ctx) bool {
		// return c.IP() == "127.0.0.1"
		// },
		Max:        max,
		Expiration: time.Duration(sec) * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return string(c.IP()) + key
		},
		LimitReached: func(c *fiber.Ctx) error {
			return helpers.TooManyRequestResponse(c)
		},
	})
}

// HTTP Throttle by user's IP address
//
//	param	key string
//	param	max int
//	param	sec int
//	receiver c *fiber.Ctx
//	return error
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
			return helpers.TooManyRequestResponse(c)
		},
	})
}

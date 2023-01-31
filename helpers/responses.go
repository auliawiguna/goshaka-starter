package helpers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Handle http success response
//
//	param c *fiber.Ctx
//	param data interface{}
//	param message string
//	return message string
func SuccessResponse(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(http.StatusAccepted).JSON(fiber.Map{
		"error":   false,
		"data":    data,
		"message": message,
	})
}

// Handle http unprocessable response
//
//	param c *fiber.Ctx
//	param data interface{}
//	param message string
//	return message string
func UnprocessableResponse(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
		"error":   true,
		"data":    data,
		"message": message,
	})
}

// Handle http unauthorised response
//
//	param c *fiber.Ctx
//	param data interface{}
//	param message string
//	return message string
func UnauthorisedResponse(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
		"error":   true,
		"data":    data,
		"message": message,
	})
}

// Handle http 404 response
//
//	param c *fiber.Ctx
//	param data interface{}
//	param message string
//	return message string
func NotFoundResponse(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(http.StatusNotFound).JSON(fiber.Map{
		"error":   true,
		"message": message,
	})
}

// Handle http 429 response
//
//	param c *fiber.Ctx
//	param data interface{}
//	param message string
//	return message string
func TooManyRequestResponse(c *fiber.Ctx) error {
	return c.Status(http.StatusTooManyRequests).JSON(fiber.Map{
		"error":   true,
		"message": "too many request",
	})
}

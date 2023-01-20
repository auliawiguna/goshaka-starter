package v1

import (
	"goshaka/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

// @Summary Say hi
// @Description Show greeting
// @Tags Root
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/v1 [get]
func IndexRoute(router fiber.Router) {
	route := router.Group("/")
	route.Get("/", func(c *fiber.Ctx) error {
		res := map[string]interface{}{
			"data": "Hi there ^^",
		}
		if err := c.JSON(res); err != nil {
			return err
		}
		return nil
	})
}

// @Security BearerAuth
// @Summary Say hi
// @Description Show greeting
// @Tags Root
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/protected [get]
func IndexProtectedRoute(router fiber.Router) {
	route := router.Group("/protected")
	route.Get("/", middlewares.ValidateJWT, func(c *fiber.Ctx) error {
		res := map[string]interface{}{
			"data": "Hi there ^^",
		}
		if err := c.JSON(res); err != nil {
			return err
		}
		return nil
	})
}

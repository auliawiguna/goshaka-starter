package controller_v1

import (
	repositories_v1 "goshaka/app/repositories"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// @Summary Login
// @Description Login
// @Tags Auth
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	loginRequest	body	structs.Login	true	"email"
// @Router /api/v1/auth/login [post]
func Login(c *fiber.Ctx) error {
	user, jwt, err := repositories_v1.Login(c)

	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": true,
			"user":  user,
			"data":  err,
		})
	}
	c.Cookie(&fiber.Cookie{
		Name:  "token",
		Value: jwt,
	})
	c.Cookie(&fiber.Cookie{
		Name:  "user_id",
		Value: strconv.FormatUint(uint64(user.ID), 10),
	})

	return c.Status(200).JSON(fiber.Map{
		"error":        false,
		"user":         user,
		"access_token": jwt,
	})
}

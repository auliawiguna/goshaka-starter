package controller_v1

import (
	"fmt"
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

// @Summary Register new account
// @Description Register new account
// @Tags Auth
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	loginRequest	body	structs.UserCreate	true	"email"
// @Router /api/v1/auth/register [post]
func Register(c *fiber.Ctx) error {
	user, err := repositories_v1.Register(c)

	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": true,
			"data":  err,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"error": false,
		"user":  user,
	})
}

// @Summary Validate registration
// @Description Validate registration
// @Tags Auth
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	loginRequest	body	structs.RegistrationToken	true	"token"
// @Router /api/v1/auth/validate-registration [post]
func ValidateRegistration(c *fiber.Ctx) error {
	user, jwt, err := repositories_v1.ValidateRegistration(c)

	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": true,
			"data":  err,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"error":        false,
		"user":         user,
		"access_token": jwt,
	})
}

// @Security BearerAuth
// @Summary My Profile
// @Description My Profile
// @Tags Auth
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/auth/my-profile [get]
func MyProfile(c *fiber.Ctx) error {
	userId := c.Locals("user_id")
	user := repositories_v1.UserShow(fmt.Sprintf("%f", userId))

	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": true,
			"data":  nil,
		})
	}
	return c.Status(404).JSON(fiber.Map{
		"error": false,
		"data":  user,
	})

}

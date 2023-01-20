package middlewares

import (
	"fmt"
	appConfig "goshaka/configs"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func ValidateJWT(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")

	secret := []byte(appConfig.GetEnv("JWT_KEY"))
	signingMethod := jwt.SigningMethodHS256

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if t.Method != signingMethod {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secret, nil
	})

	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"data":  "Unauthorised",
		})
	}

	if !token.Valid {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"data":  "Invalid Token",
		})
	}

	return c.Next()
}

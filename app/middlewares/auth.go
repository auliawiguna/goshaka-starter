package middlewares

import (
	"fmt"
	"goshaka/app/models"
	appConfig "goshaka/configs"
	"goshaka/database"
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

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Locals("user_id", claims["id"])
		c.Locals("email", claims["email"])
	}

	if !token.Valid {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"data":  "Invalid Token",
		})
	}

	return c.Next()
}

func RoleAuth(roles []string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		userId := c.Locals("user_id")

		db := database.DB
		var roleUser []*models.RoleUser

		db.Find(&roleUser, "user_id = ?", userId)

		var rolesIds []uint
		for _, role := range roleUser {
			rolesIds = append(rolesIds, role.RoleId)
		}

		var roleArray []models.Role
		db.Table("roles").Select("roles.*").Where("roles.id IN (?)", rolesIds).Where("name IN (?)", roles).Scan(&roleArray)

		if roleArray == nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"data":  "Insufficient Permission",
			})
		}

		return c.Next()
	}
}

func PermissionAuth(permissions []string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		userId := c.Locals("user_id")

		db := database.DB
		var roleUser []*models.RoleUser
		var roles []uint

		db.Find(&roleUser, "user_id = ?", userId)
		for _, role := range roleUser {
			roles = append(roles, role.RoleId)
		}

		if roles == nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"data":  "Unauthorised Action",
			})
		}
		var permissionArray []models.Permission
		db.Table("permissions").Select("permissions.*").Joins("JOIN permission_role ON permission_role.permission_id = permissions.id").Where("permission_role.role_id IN (?)", roles).Where("permissions.name IN (?)", permissions).Scan(&permissionArray)

		if permissionArray == nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"data":  "Insufficient Permission",
			})
		}

		return c.Next()
	}
}

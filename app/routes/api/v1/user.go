package v1

import (
	controllerV1 "goshaka/app/controllers/v1"
	"goshaka/app/middlewares"

	"goshaka/app/validator"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(router fiber.Router) {
	note := router.Group("/users")

	note.Get("/", middlewares.ValidateJWT, middlewares.PermissionAuth([]string{"user-read"}), controllerV1.UserIndex)
	note.Get("/:id", middlewares.ValidateJWT, middlewares.PermissionAuth([]string{"user-read"}), controllerV1.UserShow)
	note.Post("/", middlewares.ValidateJWT, middlewares.RoleAuth([]string{"admin"}), validator.CreateUserValidator, controllerV1.UserStore)
	note.Put("/:id", middlewares.ValidateJWT, middlewares.RoleAuth([]string{"admin"}), validator.CreateUserValidator, controllerV1.UserUpdate)
	note.Delete("/:id", middlewares.ValidateJWT, middlewares.RoleAuth([]string{"admin"}), controllerV1.UserDestroy)
}

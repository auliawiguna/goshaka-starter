package v1

import (
	controllerV1 "goshaka/app/controllers/v1"
	"goshaka/app/middlewares"

	"goshaka/app/validator"

	"github.com/gofiber/fiber/v2"
)

func PermissionRoute(router fiber.Router) {
	note := router.Group("/permissions")

	note.Get("/", middlewares.ValidateJWT, middlewares.PermissionAuth([]string{"permission-read"}), controllerV1.PermissionIndex)
	note.Get("/:id", middlewares.ValidateJWT, middlewares.PermissionAuth([]string{"permission-read"}), controllerV1.PermissionShow)
	note.Post("/", middlewares.ValidateJWT, middlewares.PermissionAuth([]string{"permission-create"}), validator.CreatePermissionValidator, controllerV1.PermissionStore)
	note.Put("/:id", middlewares.ValidateJWT, middlewares.PermissionAuth([]string{"permission-update"}), controllerV1.PermissionUpdate)
	note.Delete("/:id", middlewares.ValidateJWT, middlewares.PermissionAuth([]string{"permission-delete"}), controllerV1.PermissionDestroy)
}

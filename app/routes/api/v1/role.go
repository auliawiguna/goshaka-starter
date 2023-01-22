package v1

import (
	controllerV1 "goshaka/app/controllers/v1"
	"goshaka/app/middlewares"

	"goshaka/app/validator"

	"github.com/gofiber/fiber/v2"
)

func RoleRoute(router fiber.Router) {
	note := router.Group("/roles")

	note.Get("/", middlewares.ValidateJWT, middlewares.PermissionAuth([]string{"role-read"}), controllerV1.RoleIndex)
	note.Get("/:id", middlewares.ValidateJWT, middlewares.PermissionAuth([]string{"role-read"}), controllerV1.RoleShow)
	note.Post("/", middlewares.ValidateJWT, middlewares.PermissionAuth([]string{"role-create"}), validator.CreateRoleValidator, controllerV1.RoleStore)
	note.Put("/:id", middlewares.ValidateJWT, middlewares.PermissionAuth([]string{"role-update"}), controllerV1.RoleUpdate)
	note.Delete("/:id", middlewares.ValidateJWT, middlewares.PermissionAuth([]string{"role-delete"}), controllerV1.RoleDestroy)
}

package v1

import (
	controllerV1 "goshaka/app/controllers/v1"
	"goshaka/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

// Handle authorisation routes
//
//	param router fiber.ROuter
//	return	void
func FilesRoute(router fiber.Router) {
	auth := router.Group("/files")

	auth.Post("/upload", middlewares.ValidateJWT, middlewares.ThrottleByKeyAndIP("upload-file", 60, 60), controllerV1.UploadFile)
}

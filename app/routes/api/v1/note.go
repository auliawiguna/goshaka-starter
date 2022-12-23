package v1

import (
	controllerV1 "goshaka/app/controllers/v1"

	"github.com/gofiber/fiber/v2"
)

func NoteRoute(router fiber.Router) {
	note := router.Group("/notes")

	note.Get("/", controllerV1.Index)
	note.Get("/:id", controllerV1.Show)
}

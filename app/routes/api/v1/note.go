package v1

import (
	controllerV1 "goshaka/app/controllers/v1"

	"goshaka/app/validator"

	"github.com/gofiber/fiber/v2"
)

func NoteRoute(router fiber.Router) {
	note := router.Group("/notes")

	note.Get("/", controllerV1.NoteIndex)
	note.Get("/:id", controllerV1.NoteShow)
	note.Post("/", validator.CreateNoteValidator, controllerV1.NoteStore)
	note.Put("/:id", controllerV1.NoteUpdate)
	note.Delete("/:id", controllerV1.NoteDestroy)
}

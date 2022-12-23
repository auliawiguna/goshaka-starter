package repositories_v1

import (
	"goshaka/app/models"
	"goshaka/database"

	"github.com/gofiber/fiber/v2"
)

func ShowAll() ([]models.Note, bool) {
	db := database.DB
	var notes []models.Note
	var error bool = false

	db.Find(&notes)
	if len(notes) == 0 {
		error = true
	}

	return notes, error
}

func Show(id string) models.Note {
	db := database.DB
	var note models.Note

	db.Find(&note, "id = ?", id)

	return note
}

func Create(c *fiber.Ctx) (models.Note, error) {
	db := database.DB
	note := new(models.Note)

	err := c.BodyParser(note)

	if err != nil {
		return *note, err
	}

	err = db.Create(&note).Error
	return *note, err
}

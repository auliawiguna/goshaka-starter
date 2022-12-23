package repositories_v1

import (
	"goshaka/app/models"
	"goshaka/database"
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

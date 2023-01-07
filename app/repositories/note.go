package repositories_v1

import (
	"fmt"
	"goshaka/app/models"
	"goshaka/app/models/scopes"
	"goshaka/database"
	"goshaka/helpers"

	"github.com/gofiber/fiber/v2"
)

func NoteShowAll(pagination helpers.Pagination) (*helpers.Pagination, bool) {
	db := database.DB
	var notes []*models.Note
	var error bool = false

	db.Scopes(scopes.Paginate(notes, &pagination, db)).Find(&notes)
	pagination.Rows = notes

	// db.Find(&notes)
	// if len(notes) == 0 {
	// 	error = true
	// }

	return &pagination, error
}

func NoteShow(id string) models.Note {
	db := database.DB
	var note models.Note

	db.Find(&note, "id = ?", id)

	return note
}

func NoteCreate(c *fiber.Ctx) (models.Note, error) {
	db := database.DB
	note := new(models.Note)

	err := c.BodyParser(note)

	if err != nil {
		return *note, err
	}

	err = db.Create(&note).Error
	return *note, err
}

func NoteUpdate(c *fiber.Ctx, id string) (models.Note, error) {
	db := database.DB
	var note models.Note

	db.Find(&note, "id = ?", id)

	if note.ID == 0 {
		return note, fmt.Errorf("not found")
	}
	noteUpdate := new(models.Note)

	err := c.BodyParser(noteUpdate)

	if err != nil {
		return *noteUpdate, err
	}

	db.Model(&note).Where("id = ?", id).UpdateColumns(&noteUpdate)

	return note, err
}

func NoteDestroy(c *fiber.Ctx, id string) (models.Note, error) {
	db := database.DB
	var note models.Note

	db.Find(&note, "id = ?", id)

	if note.ID == 0 {
		return note, fmt.Errorf("not found")
	}

	//To soft delete, just remove .Unscoped()
	err := db.Unscoped().Delete(&note).Error

	return note, err
}

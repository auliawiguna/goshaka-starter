package controller_v1

import (
	repositories_v1 "goshaka/app/repositories"
	"goshaka/helpers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func NoteIndex(c *fiber.Ctx) error {
	var pagination helpers.Pagination
	pagination.Limit, _ = strconv.Atoi(c.Query("limit"))
	pagination.Page, _ = strconv.Atoi(c.Query("page"))
	pagination.Sort = c.Query("sort")
	notes, err := repositories_v1.NoteShowAll(pagination)

	return c.JSON(fiber.Map{
		"error": err,
		"data":  notes,
	})
}

func NoteShow(c *fiber.Ctx) error {
	note := repositories_v1.NoteShow(c.Params("id"))

	if note.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": true,
			"data":  nil,
		})
	}
	return c.Status(404).JSON(fiber.Map{
		"error": false,
		"data":  note,
	})
}

func NoteStore(c *fiber.Ctx) error {
	note, err := repositories_v1.NoteCreate(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": true,
			"data":  err,
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"data":  note,
	})
}

func NoteUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	note, err := repositories_v1.NoteUpdate(c, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": true,
			"data":  err,
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"data":  note,
	})
}

func NoteDestroy(c *fiber.Ctx) error {
	id := c.Params("id")
	note, err := repositories_v1.NoteDestroy(c, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": true,
			"data":  err,
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"data":  note,
	})
}

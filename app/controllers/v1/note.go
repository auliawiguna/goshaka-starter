package controller_v1

import (
	repositories_v1 "goshaka/app/repositories"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	notes, err := repositories_v1.ShowAll()

	if len(notes) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": err,
			"data":  nil,
		})
	}

	return c.JSON(fiber.Map{
		"error": err,
		"data":  notes,
	})
}

func Show(c *fiber.Ctx) error {
	note := repositories_v1.Show(c.Params("id"))

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

func Store(c *fiber.Ctx) error {
	note, err := repositories_v1.Create(c)
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

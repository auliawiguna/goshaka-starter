package controller_v1

import (
	repositories_v1 "goshaka/app/repositories"
	"goshaka/helpers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// @Summary Show notes
// @Description Show notes
// @Tags Notes
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	limit	query	int	false	"Default 10"	default(10)
// @Param	page	query	int	false	"Default 10"	default(1)
// @Param	sort	query	string	false	"Sorting"	Enums(ID asc, ID desc, title asc, title desc)
// @Router /api/v1/notes [get]
func NoteIndex(c *fiber.Ctx) error {
	var pagination helpers.Pagination
	pagination.Limit, _ = strconv.Atoi(c.Query("limit"))
	pagination.Page, _ = strconv.Atoi(c.Query("page"))
	pagination.Sort = c.Query("sort")
	notes := repositories_v1.NoteShowAll(pagination)

	return helpers.SuccessResponse(c, notes, "success")
}

// @Summary Show detail note
// @Description Show detail note
// @Tags Notes
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	id	path	int	false	"ID"	default(10)
// @Router /api/v1/notes/{id} [get]
func NoteShow(c *fiber.Ctx) error {
	note := repositories_v1.NoteShow(c.Params("id"))

	if note.ID == 0 {
		return helpers.NotFoundResponse(c, note, "not found")
	}
	return helpers.SuccessResponse(c, note, "success")
}

// @Summary Create new note
// @Description Create new note
// @Tags Notes
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	notesRequest	body	structs.NoteCreate	true	"title"
// @Router /api/v1/notes [post]
func NoteStore(c *fiber.Ctx) error {
	note, err := repositories_v1.NoteCreate(c)
	if err != nil {
		return helpers.UnprocessableResponse(c, note, err.Error())
	}
	return helpers.SuccessResponse(c, note, "success")
}

// @Summary Update existing note
// @Description Update existing note
// @Tags Notes
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	id	path	int	false	"ID"	default(10)
// @Param	notesRequest	body	structs.NoteCreate	true	"title"
// @Router /api/v1/notes/{id} [put]
func NoteUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	note, err := repositories_v1.NoteUpdate(c, id)
	if err != nil {
		return helpers.UnprocessableResponse(c, note, err.Error())
	}
	return helpers.SuccessResponse(c, note, "success")
}

// @Summary Delete existing note
// @Description Delete existing note
// @Tags Notes
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	id	path	int	false	"ID"	default(10)
// @Router /api/v1/notes/{id} [delete]
func NoteDestroy(c *fiber.Ctx) error {
	id := c.Params("id")
	note, err := repositories_v1.NoteDestroy(c, id)
	if err != nil {
		return helpers.UnprocessableResponse(c, note, err.Error())
	}
	return helpers.SuccessResponse(c, note, "success")
}

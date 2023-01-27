package controller_v1

import (
	repositories_v1 "goshaka/app/repositories"
	"goshaka/helpers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// @Security BearerAuth
// @Summary Show users
// @Description Show users
// @Tags Users
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	limit	query	int	false	"Default 10"	default(10)
// @Param	page	query	int	false	"Default 10"	default(1)
// @Param	sort	query	string	false	"Sorting"	Enums(ID asc, ID desc, title asc, title desc)
// @Router /api/v1/users [get]
func UserIndex(c *fiber.Ctx) error {
	var pagination helpers.Pagination
	pagination.Limit, _ = strconv.Atoi(c.Query("limit"))
	pagination.Page, _ = strconv.Atoi(c.Query("page"))
	pagination.Sort = c.Query("sort")
	users, err := repositories_v1.UserShowAll(pagination)

	return c.JSON(fiber.Map{
		"error": err,
		"data":  users,
	})
}

// @Security BearerAuth
// @Summary Show detail user
// @Description Show detail user
// @Tags Users
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	id	path	int	false	"ID"	default(10)
// @Router /api/v1/users/{id} [get]
func UserShow(c *fiber.Ctx) error {
	user := repositories_v1.UserShow(c.Params("id"))

	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": true,
			"data":  nil,
		})
	}
	return c.Status(404).JSON(fiber.Map{
		"error": false,
		"data":  user,
	})
}

// @Security BearerAuth
// @Summary Create new user
// @Description Create new user
// @Tags Users
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	usersRequest	body	structs.UserCreate	true	"title"
// @Router /api/v1/users [post]
func UserStore(c *fiber.Ctx) error {
	user, err := repositories_v1.UserCreate(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": true,
			"data":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"data":  user,
	})
}

// @Security BearerAuth
// @Summary Update existing user
// @Description Update existing user
// @Tags Users
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	id	path	int	false	"ID"	default(10)
// @Param	usersRequest	body	structs.UserCreate	true	"title"
// @Router /api/v1/users/{id} [put]
func UserUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := repositories_v1.UserUpdate(c, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": true,
			"data":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"data":  user,
	})
}

// @Security BearerAuth
// @Summary Delete existing user
// @Description Delete existing user
// @Tags Users
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	id	path	int	false	"ID"	default(10)
// @Router /api/v1/users/{id} [delete]
func UserDestroy(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := repositories_v1.UserDestroy(c, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": true,
			"data":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"data":  user,
	})
}

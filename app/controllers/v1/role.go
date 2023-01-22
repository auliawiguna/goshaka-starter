package controller_v1

import (
	repositories_v1 "goshaka/app/repositories"
	"goshaka/helpers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// @Security BearerAuth
// @Summary Show roles
// @Description Show roles
// @Tags Roles
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	limit	query	int	false	"Default 10"	default(10)
// @Param	page	query	int	false	"Default 10"	default(1)
// @Param	sort	query	string	false	"Sorting"	Enums(ID asc, ID desc, title asc, title desc)
// @Router /api/v1/roles [get]
func RoleIndex(c *fiber.Ctx) error {
	var pagination helpers.Pagination
	pagination.Limit, _ = strconv.Atoi(c.Query("limit"))
	pagination.Page, _ = strconv.Atoi(c.Query("page"))
	pagination.Sort = c.Query("sort")
	roles, err := repositories_v1.RoleShowAll(pagination)

	return c.JSON(fiber.Map{
		"error": err,
		"data":  roles,
	})
}

// @Security BearerAuth
// @Summary Show detail role
// @Description Show detail role
// @Tags Roles
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	id	path	int	false	"ID"	default(10)
// @Router /api/v1/roles/{id} [get]
func RoleShow(c *fiber.Ctx) error {
	role := repositories_v1.RoleShow(c.Params("id"))

	if role.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": true,
			"data":  nil,
		})
	}
	return c.Status(404).JSON(fiber.Map{
		"error": false,
		"data":  role,
	})
}

// @Security BearerAuth
// @Summary Create new role
// @Description Create new role
// @Tags Roles
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	rolesRequest	body	structs.RoleCreate	true	"title"
// @Router /api/v1/roles [post]
func RoleStore(c *fiber.Ctx) error {
	role, err := repositories_v1.RoleCreate(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": true,
			"data":  err,
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"data":  role,
	})
}

// @Security BearerAuth
// @Summary Update existing role
// @Description Update existing role
// @Tags Roles
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	id	path	int	false	"ID"	default(10)
// @Param	rolesRequest	body	structs.RoleCreate	true	"title"
// @Router /api/v1/roles/{id} [put]
func RoleUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	role, err := repositories_v1.RoleUpdate(c, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": true,
			"data":  err,
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"data":  role,
	})
}

// @Security BearerAuth
// @Summary Delete existing role
// @Description Delete existing role
// @Tags Roles
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	id	path	int	false	"ID"	default(10)
// @Router /api/v1/roles/{id} [delete]
func RoleDestroy(c *fiber.Ctx) error {
	id := c.Params("id")
	role, err := repositories_v1.RoleDestroy(c, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": true,
			"data":  err,
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"data":  role,
	})
}

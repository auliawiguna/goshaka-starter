package controller_v1

import (
	repositories_v1 "goshaka/app/repositories"
	"goshaka/helpers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// @Security BearerAuth
// @Summary Show permissions
// @Description Show permissions
// @Tags Permissions
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	limit	query	int	false	"Default 10"	default(10)
// @Param	page	query	int	false	"Default 10"	default(1)
// @Param	sort	query	string	false	"Sorting"	Enums(ID asc, ID desc, title asc, title desc)
// @Router /api/v1/permissions [get]
func PermissionIndex(c *fiber.Ctx) error {
	var pagination helpers.Pagination
	pagination.Limit, _ = strconv.Atoi(c.Query("limit"))
	pagination.Page, _ = strconv.Atoi(c.Query("page"))
	pagination.Sort = c.Query("sort")
	permissions, err := repositories_v1.PermissionShowAll(pagination)

	return c.JSON(fiber.Map{
		"error": err,
		"data":  permissions,
	})
}

// @Security BearerAuth
// @Summary Show detail permission
// @Description Show detail permission
// @Tags Permissions
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	id	path	int	false	"ID"	default(10)
// @Router /api/v1/permissions/{id} [get]
func PermissionShow(c *fiber.Ctx) error {
	permission := repositories_v1.PermissionShow(c.Params("id"))

	if permission.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": true,
			"data":  nil,
		})
	}
	return c.Status(404).JSON(fiber.Map{
		"error": false,
		"data":  permission,
	})
}

// @Security BearerAuth
// @Summary Create new permission
// @Description Create new permission
// @Tags Permissions
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	permissionsRequest	body	structs.PermissionCreate	true	"title"
// @Router /api/v1/permissions [post]
func PermissionStore(c *fiber.Ctx) error {
	permission, err := repositories_v1.PermissionCreate(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": true,
			"data":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"data":  permission,
	})
}

// @Security BearerAuth
// @Summary Update existing permission
// @Description Update existing permission
// @Tags Permissions
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	id	path	int	false	"ID"	default(10)
// @Param	permissionsRequest	body	structs.PermissionCreate	true	"title"
// @Router /api/v1/permissions/{id} [put]
func PermissionUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	permission, err := repositories_v1.PermissionUpdate(c, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": true,
			"data":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"data":  permission,
	})
}

// @Security BearerAuth
// @Summary Delete existing permission
// @Description Delete existing permission
// @Tags Permissions
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	id	path	int	false	"ID"	default(10)
// @Router /api/v1/permissions/{id} [delete]
func PermissionDestroy(c *fiber.Ctx) error {
	id := c.Params("id")
	permission, err := repositories_v1.PermissionDestroy(c, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": true,
			"data":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"data":  permission,
	})
}

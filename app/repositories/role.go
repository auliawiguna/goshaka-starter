package repositories_v1

import (
	"fmt"
	"goshaka/app/models"
	"goshaka/app/models/scopes"
	"goshaka/database"
	"goshaka/helpers"

	"github.com/gofiber/fiber/v2"
)

func RoleShowAll(pagination helpers.Pagination) (*helpers.Pagination, bool) {
	db := database.DB
	var roles []*models.Role
	var error bool = false

	db.Scopes(scopes.Paginate(roles, &pagination, db)).Find(&roles)
	pagination.Rows = roles

	return &pagination, error
}

func RoleShow(id string) models.Role {
	db := database.DB
	var role models.Role

	db.Find(&role, "id = ?", id)

	return role
}

func RoleCreate(c *fiber.Ctx) (models.Role, error) {
	db := database.DB
	role := new(models.Role)

	err := c.BodyParser(role)

	if err != nil {
		return *role, err
	}

	err = db.Create(&role).Error
	return *role, err
}

func RoleUpdate(c *fiber.Ctx, id string) (models.Role, error) {
	db := database.DB
	var role models.Role

	db.Find(&role, "id = ?", id)

	if role.ID == 0 {
		return role, fmt.Errorf("not found")
	}
	roleUpdate := new(models.Role)

	err := c.BodyParser(roleUpdate)

	if err != nil {
		return *roleUpdate, err
	}

	db.Model(&role).Where("id = ?", id).UpdateColumns(&roleUpdate)

	return role, err
}

func RoleDestroy(c *fiber.Ctx, id string) (models.Role, error) {
	db := database.DB
	var role models.Role

	db.Find(&role, "id = ?", id)

	if role.ID == 0 {
		return role, fmt.Errorf("not found")
	}

	//To soft delete, just remove .Unscoped()
	err := db.Unscoped().Delete(&role).Error

	return role, err
}

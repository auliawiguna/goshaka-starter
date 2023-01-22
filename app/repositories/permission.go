package repositories_v1

import (
	"fmt"
	"goshaka/app/models"
	"goshaka/app/models/scopes"
	"goshaka/database"
	"goshaka/helpers"

	"github.com/gofiber/fiber/v2"
)

func PermissionShowAll(pagination helpers.Pagination) (*helpers.Pagination, bool) {
	db := database.DB
	var permissions []*models.Permission
	var error bool = false

	db.Scopes(scopes.Paginate(permissions, &pagination, db)).Find(&permissions)
	pagination.Rows = permissions

	return &pagination, error
}

func PermissionShow(id string) models.Permission {
	db := database.DB
	var permission models.Permission

	db.Find(&permission, "id = ?", id)

	return permission
}

func PermissionCreate(c *fiber.Ctx) (models.Permission, error) {
	db := database.DB
	permission := new(models.Permission)

	err := c.BodyParser(permission)

	if err != nil {
		return *permission, err
	}

	err = db.Create(&permission).Error
	return *permission, err
}

func PermissionUpdate(c *fiber.Ctx, id string) (models.Permission, error) {
	db := database.DB
	var permission models.Permission

	db.Find(&permission, "id = ?", id)

	if permission.ID == 0 {
		return permission, fmt.Errorf("not found")
	}
	permissionUpdate := new(models.Permission)

	err := c.BodyParser(permissionUpdate)

	if err != nil {
		return *permissionUpdate, err
	}

	db.Model(&permission).Where("id = ?", id).UpdateColumns(&permissionUpdate)

	return permission, err
}

func PermissionDestroy(c *fiber.Ctx, id string) (models.Permission, error) {
	db := database.DB
	var permission models.Permission

	db.Find(&permission, "id = ?", id)

	if permission.ID == 0 {
		return permission, fmt.Errorf("not found")
	}

	//To soft delete, just remove .Unscoped()
	err := db.Unscoped().Delete(&permission).Error

	return permission, err
}

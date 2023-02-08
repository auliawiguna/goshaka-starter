package repositories_v1

import (
	"fmt"
	"goshaka/app/models"
	"goshaka/app/models/scopes"
	"goshaka/database"
	"goshaka/helpers"

	"github.com/gofiber/fiber/v2"
)

func FileShowAll(pagination helpers.Pagination) *helpers.Pagination {
	db := database.DB
	var files []*models.File

	db.Scopes(scopes.Paginate(files, &pagination, db)).Find(&files)
	pagination.Rows = files

	return &pagination
}

func FileShow(id string) models.File {
	db := database.DB
	var file models.File

	db.Find(&file, "id = ?", id)

	return file
}

func FileCreate(arr interface{}) (models.File, error) {
	db := database.DB

	file := &models.File{}
	err := db.Model(file).Create(arr).Error
	var newFile models.File
	db.Last(&newFile)

	return newFile, err
}

func FileUpdate(c *fiber.Ctx, id string) (models.File, error) {
	db := database.DB
	var file models.File

	db.Find(&file, "id = ?", id)

	if file.ID == 0 {
		return file, fmt.Errorf("not found")
	}
	fileUpdate := new(models.File)

	err := c.BodyParser(fileUpdate)

	if err != nil {
		return *fileUpdate, err
	}

	db.Model(&file).Where("id = ?", id).UpdateColumns(&fileUpdate)

	return file, err
}

func FileDestroy(c *fiber.Ctx, id string) (models.File, error) {
	db := database.DB
	var file models.File

	db.Find(&file, "id = ?", id)

	if file.ID == 0 {
		return file, fmt.Errorf("not found")
	}

	// To soft delete, just remove .Unscoped()
	err := db.Unscoped().Delete(&file).Error

	return file, err
}

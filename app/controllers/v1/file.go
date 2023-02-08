package controller_v1

import (
	"fmt"
	repositories_v1 "goshaka/app/repositories"
	"goshaka/helpers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// @Security BearerAuth
// @Summary Upload a file to AWS S3
// @Description Upload a file to a specified AWS S3 bucket
// @Accept multipart/form-data
// @Tags File Upload
// @Produce json
// @Param file formData file true "File to upload"
// @Param string formData string true "A string value (optional)"
// @Success 200 {object} string "Success message"
// @Failure 400 {object} string "Error message"
// @Failure 500 {object} string "Error message"
// @Router /api/v1/files/upload [post]
func UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")

	if err != nil {
		return helpers.UnprocessableResponse(c, file, err.Error())
	}

	// Upload the file to S3
	upload, err := helpers.UploadFileToS3(file, "/files/")
	if err != nil {
		return helpers.UnprocessableResponse(c, file, err.Error())
	}

	return helpers.SuccessResponse(c, upload, "success")
}

// @Security BearerAuth
// @Summary Upload a file to AWS S3
// @Description Upload a file to a specified AWS S3 bucket
// @Accept multipart/form-data
// @Tags File Upload
// @Produce json
// @Param file formData file true "File to upload"
// @Param string formData string true "A string value (optional)"
// @Success 200 {object} string "Success message"
// @Failure 400 {object} string "Error message"
// @Failure 500 {object} string "Error message"
// @Router /api/v1/files/userfile [post]
func UploadUserFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	uid := c.Locals("user_id")

	if err != nil {
		return helpers.UnprocessableResponse(c, file, err.Error())
	}

	// Upload the file to S3
	upload, err := helpers.UploadFileToS3(file, "/avatar/")
	if err != nil {
		return helpers.UnprocessableResponse(c, file, err.Error())
	}

	uploadMap, ok := upload.(map[string]interface{})
	if !ok {
		return helpers.UnprocessableResponse(c, file, "cannot parse aws file")
	}

	f := make(map[string]interface{})
	f["path"] = uploadMap["AWSUrl"]
	f["user_id"] = uid
	f["filename"] = uploadMap["filename"]
	f["mimetype"] = uploadMap["mimetype"]
	f["size"] = uploadMap["size"]

	nf, errB := repositories_v1.FileCreate(f)
	if errB != nil {
		return helpers.UnprocessableResponse(c, file, errB.Error())
	}

	return helpers.SuccessResponse(c, nf, "success")
}

// @Security BearerAuth
// @Summary Show notes
// @Description Show notes
// @Tags File Upload
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	limit	query	int	false	"Default 10"	default(10)
// @Param	page	query	int	false	"Default 10"	default(1)
// @Param	sort	query	string	false	"Sorting"	Enums(ID asc, ID desc, title asc, title desc)
// @Router /api/v1/files/userfile [get]
func GetUserFiles(c *fiber.Ctx) error {
	uid := fmt.Sprintf("%v", c.Locals("user_id"))

	var pagination helpers.Pagination
	pagination.Limit, _ = strconv.Atoi(c.Query("limit"))
	pagination.Page, _ = strconv.Atoi(c.Query("page"))
	pagination.Sort = c.Query("sort")
	notes := repositories_v1.FileShowAll(pagination, uid)

	return helpers.SuccessResponse(c, notes, "success")

}

// @Security BearerAuth
// @Summary Show detail file
// @Description Show detail file
// @Tags File Upload
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	id	path	int	false	"ID"	default(10)
// @Router /api/v1/files/userfile/{id} [get]
func ShowUserFile(c *fiber.Ctx) error {
	uid := fmt.Sprintf("%v", c.Locals("user_id"))

	note := repositories_v1.FileShow(c.Params("id"), uid)

	if note.ID == 0 {
		return helpers.NotFoundResponse(c, note, "not found")
	}
	return helpers.SuccessResponse(c, note, "success")
}

// @Security BearerAuth
// @Summary Show detail file
// @Description Show detail file
// @Tags File Upload
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	id	path	int	false	"ID"	default(10)
// @Router /api/v1/files/userfile/{id} [delete]
func DeleteUserFile(c *fiber.Ctx) error {
	uid := fmt.Sprintf("%v", c.Locals("user_id"))

	note, err := repositories_v1.FileDestroy(c.Params("id"), uid)

	if err != nil {
		return helpers.InternalServerErorResponse(c, err.Error())
	}

	if note.ID == 0 {
		return helpers.NotFoundResponse(c, note, "not found")
	}
	return helpers.SuccessResponse(c, note, "success")
}

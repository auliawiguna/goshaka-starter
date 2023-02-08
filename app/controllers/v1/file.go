package controller_v1

import (
	repositories_v1 "goshaka/app/repositories"
	"goshaka/helpers"

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

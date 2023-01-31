package validator

import (
	"goshaka/app/structs"
	"goshaka/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CreateRoleValidator(c *fiber.Ctx) error {
	var errors []*structs.IError
	body := new(structs.RoleCreate)
	c.BodyParser(&body)

	err := Validator.Struct(body)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el structs.IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}
		return helpers.UnprocessableResponse(c, errors, "unprocessable entity")
	}

	return c.Next()
}

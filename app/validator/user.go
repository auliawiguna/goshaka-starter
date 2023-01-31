package validator

import (
	"goshaka/app/structs"
	"goshaka/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Validate payload on create new user
//
//	param c *fiber.Ctx
//	return error
func CreateUserValidator(c *fiber.Ctx) error {
	var errors []*structs.IError
	body := new(structs.UserCreate)
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

// Validate payload on register a new user
//
//	param c *fiber.Ctx
//	return error
func RegistrationValidator(c *fiber.Ctx) error {
	var errors []*structs.IError
	body := new(structs.RegistrationToken)
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

// Validate payload on request verification token
//
//	param c *fiber.Ctx
//	return error
func ResendTokenValidator(c *fiber.Ctx) error {
	var errors []*structs.IError
	body := new(structs.ResendToken)
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

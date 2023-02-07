package validator

import (
	"fmt"
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
	errb := c.BodyParser(&body)
	if errb != nil {
		return fmt.Errorf("%s", "error parsing data")
	}

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
	errb := c.BodyParser(&body)
	if errb != nil {
		return fmt.Errorf("%s", "error parsing data")
	}

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
	errb := c.BodyParser(&body)
	if errb != nil {
		return fmt.Errorf("%s", "error parsing data")
	}

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

// Validate payload on update profile
//
//	param c *fiber.Ctx
//	return error
func ProfileUpdateValidator(c *fiber.Ctx) error {
	var errors []*structs.IError
	body := new(structs.ProfileUpdate)
	errb := c.BodyParser(&body)
	if errb != nil {
		return fmt.Errorf("%s", "error parsing data")
	}

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

// Validate payload on update email address
//
//	param c *fiber.Ctx
//	return error
func EmailUpdateValidator(c *fiber.Ctx) error {
	var errors []*structs.IError
	body := new(structs.EmailUpdate)
	errb := c.BodyParser(&body)
	if errb != nil {
		return fmt.Errorf("%s", "error parsing data")
	}

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

package validator

import (
	"fmt"
	"goshaka/app/structs"
	"goshaka/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func LoginValidator(c *fiber.Ctx) error {
	var errors []*structs.IError
	body := new(structs.Login)
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

func RequestResetPasswordValidator(c *fiber.Ctx) error {
	var errors []*structs.IError
	body := new(structs.RequestResetPassword)
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

func ResetPasswordValidator(c *fiber.Ctx) error {
	var errors []*structs.IError
	body := new(structs.ResetPassword)
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

func GoogleOneTap(c *fiber.Ctx) error {
	var errors []*structs.IError
	body := new(structs.GoogleOneTap)
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

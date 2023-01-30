package validator

import (
	"goshaka/app/structs"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func LoginValidator(c *fiber.Ctx) error {
	var errors []*structs.IError
	body := new(structs.Login)
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
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.Next()
}

func RequestResetPasswordValidator(c *fiber.Ctx) error {
	var errors []*structs.IError
	body := new(structs.RequestResetPassword)
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
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.Next()
}

func ResetPasswordValidator(c *fiber.Ctx) error {
	var errors []*structs.IError
	body := new(structs.ResetPassword)
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
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.Next()
}

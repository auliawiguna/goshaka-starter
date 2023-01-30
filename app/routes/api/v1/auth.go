package v1

import (
	controllerV1 "goshaka/app/controllers/v1"
	"goshaka/app/middlewares"

	"goshaka/app/validator"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(router fiber.Router) {
	auth := router.Group("/auth")

	auth.Post("/login", validator.LoginValidator, controllerV1.Login)
	auth.Post("/register", validator.CreateUserValidator, controllerV1.Register)
	auth.Post("/validate-registration", validator.RegistrationValidator, controllerV1.ValidateRegistration)
	auth.Post("/request-reset-password", validator.RequestResetPasswordValidator, controllerV1.RequestResetPassword)
	auth.Post("/reset-password", validator.ResetPasswordValidator, controllerV1.RequestResetPassword)
	auth.Get("/my-profile", middlewares.ValidateJWT, controllerV1.MyProfile)
}

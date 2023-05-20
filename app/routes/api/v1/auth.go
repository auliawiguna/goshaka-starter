package v1

import (
	controllerV1 "goshaka/app/controllers/v1"
	"goshaka/app/middlewares"

	"goshaka/app/validator"

	"github.com/gofiber/fiber/v2"
)

// Handle authorisation routes
//
//	param router fiber.ROuter
//	return	void
func AuthRoute(router fiber.Router) {
	auth := router.Group("/auth")

	// Login
	auth.Post("/login", validator.LoginValidator, controllerV1.Login)
	// Registration
	auth.Post("/register", middlewares.ThrottleByKeyAndIP("register", 60, 60), validator.CreateUserValidator, controllerV1.Register)
	auth.Post("/validate-registration", validator.RegistrationValidator, controllerV1.ValidateRegistration)
	auth.Post("/resend-registration-token", middlewares.ThrottleByKeyAndIP("resend-registration-token", 60, 60), validator.ResendTokenValidator, controllerV1.ResendRegistrationToken)
	// OTP login
	auth.Post("/otp/request", middlewares.ThrottleByKeyAndIP("otp_request", 60, 60), validator.RequestOtpValidator, controllerV1.SendOtp)
	auth.Post("/otp/validate", middlewares.ThrottleByKeyAndIP("otp_validate", 60, 60), validator.RequestValidateOtpValidator, controllerV1.ValidateOtp)
	// Forgot Password
	auth.Post("/request-reset-password", middlewares.ThrottleByKeyAndIP("request-reset-password", 60, 60), validator.RequestResetPasswordValidator, controllerV1.RequestResetPassword)
	auth.Post("/reset-password", validator.ResetPasswordValidator, controllerV1.ResetPassword)
	// Handle Google One Tap login
	auth.Post("/google-one-tap", validator.GoogleOneTap, controllerV1.GoogleOneTap)
	// My Profile
	auth.Get("/my-profile", middlewares.ValidateJWT, controllerV1.MyProfile)
	auth.Put("/my-profile", middlewares.ValidateJWT, validator.ProfileUpdateValidator, controllerV1.UpdateProfile)
	auth.Post("/validate-new-email", middlewares.ValidateJWT, validator.EmailUpdateValidator, controllerV1.UpdateEmail)
}

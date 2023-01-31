package controller_v1

import (
	"encoding/json"
	"fmt"
	repositories_v1 "goshaka/app/repositories"
	"goshaka/app/structs"
	"goshaka/helpers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// @Summary Login
// @Description Login
// @Tags Auth
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	loginRequest	body	structs.Login	true	"email"
// @Router /api/v1/auth/login [post]
func Login(c *fiber.Ctx) error {
	user, jwt, err := repositories_v1.Login(c)

	if err != nil {
		return helpers.UnprocessableResponse(c, user, err.Error())
	}
	c.Cookie(&fiber.Cookie{
		Name:  "token",
		Value: jwt,
	})
	c.Cookie(&fiber.Cookie{
		Name:  "user_id",
		Value: strconv.FormatUint(uint64(user.ID), 10),
	})

	res := map[string]interface{}{
		"user":         user,
		"access_token": jwt,
	}

	return helpers.SuccessResponse(c, res, "success")
}

// @Summary Register new account
// @Description Register new account
// @Tags Auth
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	loginRequest	body	structs.UserCreate	true	"email"
// @Router /api/v1/auth/register [post]
func Register(c *fiber.Ctx) error {
	user, err := repositories_v1.Register(c)

	if err != nil {
		return helpers.UnprocessableResponse(c, err, err.Error())
	}

	return helpers.SuccessResponse(c, user, "success")
}

// @Summary Validate registration
// @Description Validate registration
// @Tags Auth
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	loginRequest	body	structs.RegistrationToken	true	"token"
// @Router /api/v1/auth/validate-registration [post]
func ValidateRegistration(c *fiber.Ctx) error {
	user, jwt, err := repositories_v1.ValidateRegistration(c)

	if err != nil {
		return helpers.UnauthorisedResponse(c, err, err.Error())
	}

	res := map[string]interface{}{
		"user":         user,
		"access_token": jwt,
	}

	return helpers.SuccessResponse(c, res, "success")
}

// @Summary Resend registration token
// @Description Resend registration token
// @Tags Auth
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	loginRequest	body	structs.ResendToken	true	"email"
// @Router /api/v1/auth/resend-registration-token [post]
func ResendRegistrationToken(c *fiber.Ctx) error {
	err := repositories_v1.ResendNewRegistrationToken(c)

	if err != nil {
		return helpers.UnprocessableResponse(c, err, err.Error())
	}

	return helpers.SuccessResponse(c, err, "success")
}

// @Summary Request reset password
// @Description Request reset password
// @Tags Auth
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	loginRequest	body	structs.RequestResetPassword	true	"email"
// @Router /api/v1/auth/request-reset-password [post]
func RequestResetPassword(c *fiber.Ctx) error {
	var requestResetPasswordStructure structs.RequestResetPassword

	body := c.Body()
	err := json.Unmarshal(body, &requestResetPasswordStructure)

	if err != nil {
		return helpers.UnprocessableResponse(c, err, err.Error())
	}

	//Throttle, 2 requests per email per 60 secs
	var rateLimit bool = helpers.RateLimit("requestResetPass"+requestResetPasswordStructure.Email, 2, 60)
	if !rateLimit {
		return helpers.TooManyRequestResponse(c)
	}

	msg, err := repositories_v1.RequestResetPassword(c)

	if err != nil {
		return helpers.UnprocessableResponse(c, err, err.Error())
	}

	return helpers.SuccessResponse(c, err, msg)
}

// @Summary Request reset password
// @Description Request reset password
// @Tags Auth
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param	loginRequest	body	structs.ResetPassword	true	"email"
// @Router /api/v1/auth/reset-password [post]
func ResetPassword(c *fiber.Ctx) error {
	msg, err := repositories_v1.ResetPassword(c)

	if err != nil {
		return helpers.UnprocessableResponse(c, err, err.Error())
	}

	return helpers.SuccessResponse(c, err, msg)
}

// @Security BearerAuth
// @Summary My Profile
// @Description My Profile
// @Tags Auth
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/auth/my-profile [get]
func MyProfile(c *fiber.Ctx) error {
	userId := c.Locals("user_id")
	user := repositories_v1.UserShow(fmt.Sprintf("%f", userId))

	if user.ID == 0 {
		return helpers.NotFoundResponse(c, user, "not found")
	}

	return helpers.SuccessResponse(c, user, "success")

}

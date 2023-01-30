package structs

type RequestResetPassword struct {
	Email string `json:"email" validate:"required"`
}

type ResetPassword struct {
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,eqfield=Password"`
	Token                string `json:"token" validate:"required"`
}

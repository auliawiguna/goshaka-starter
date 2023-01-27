package structs

type RequestResetPassword struct {
	Email string `json:"email" validate:"required"`
}

package structs

type RequestResetPassword struct {
	Email string `json:"email" validate:"required"`
}

type ResetPassword struct {
	Email  string `json:"email" validate:"required"`
	Token  string `json:"token" validate:"required"`
	Number string `json:"number" validate:"required"`
}

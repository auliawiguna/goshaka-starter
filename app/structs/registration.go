package structs

type RegistrationToken struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Token    string `json:"token" validate:"required"`
}

type ResendToken struct {
	Email string `json:"email" validate:"required"`
}

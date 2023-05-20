package structs

type Login struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type GoogleOneTap struct {
	IdToken string `json:"id_token" validate:"required"`
}

type EmailOnly struct {
	Email string `json:"email" validate:"required"`
}

type TokenOnly struct {
	Token string `json:"token" validate:"required"`
}

type EmailAndToken struct {
	Email string `json:"email" validate:"required"`
	Token string `json:"token" validate:"required"`
}

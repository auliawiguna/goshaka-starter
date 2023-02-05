package structs

type Login struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type GoogleOneTap struct {
	IdToken string `json:"id_token" validate:"required"`
}

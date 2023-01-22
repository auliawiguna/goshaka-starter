package structs

type RoleCreate struct {
	Title    string `json:"title" validate:"required"`
	SubTitle string `json:"subtitle" validate:"required"`
	Text     string `json:"text" validate:"required"`
}

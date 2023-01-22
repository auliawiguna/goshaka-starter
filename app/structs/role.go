package structs

type RoleCreate struct {
	Name    string `json:"name" validate:"required"`
	Display string `json:"display" validate:"required"`
}

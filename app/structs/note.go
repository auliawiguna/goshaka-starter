package structs

type NoteCreate struct {
	Title    string `json:"title"`
	SubTitle string `json:"subtitle"`
	Text     string `json:"text"`
}

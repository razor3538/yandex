package models

// SaveBytes модель для сохранения текста
type SaveBytes struct {
	Bytes string `json:"bytes"`
	Meta  string `json:"meta"`
	Name  string `json:"name_pair"`
}

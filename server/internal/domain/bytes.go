package domain

import "github.com/gofrs/uuid"

// Byte Структура хранящихся в базе данных бинарные данные
type Byte struct {
	Base
	UserId uuid.UUID `gorm:"type:varchar" json:"user_id"`
	Name   string    `gorm:"type:varchar" json:"name"`
	Bytes  string    `gorm:"type:varchar" json:"text"`
	Meta   string    `gorm:"type:text" json:"meta"`
}

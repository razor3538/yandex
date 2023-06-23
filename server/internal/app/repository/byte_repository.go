package repositories

import (
	"server/config"
	"server/internal/domain"
)

// ByteRepo структура
type ByteRepo struct{}

// NewByteRepo метод возвращает указатель на структуру ByteRepo со всеми ее методами
func NewByteRepo() *ByteRepo {
	return &ByteRepo{}
}

// Save сохраняет бинарные данные
func (br *ByteRepo) Save(byte domain.Byte) (domain.Byte, error) {
	var existingText domain.Text

	if err := config.DB.
		Table("bytes as b").
		Select("b.*").
		Where("b.name = ? and b.user_id = ?", byte.Name, byte.UserId).
		Scan(&existingText).
		Error; err != nil {
		if err.Error() != "record not found" {
			return domain.Byte{}, err
		}
	}

	if existingText.Name != "" {
		existingText.Text = byte.Bytes

		if err := config.DB.Save(&existingText).Error; err != nil {
			return domain.Byte{}, err
		}
		return byte, nil
	}

	if err := config.DB.
		Create(&byte).
		Error; err != nil {
		return domain.Byte{}, err
	}

	return byte, nil
}

// GetByKey возвращает текст по переданному имени
func (br *ByteRepo) GetByKey(value, userId string) (domain.Byte, error) {
	var byte domain.Byte
	err := config.DB.
		Table("bytes as b").
		Select("b.*").
		Where("b.name = ? and b.user_id = ?", value, userId).
		Scan(&byte).Error

	return byte, err
}

// Delete удаляет текст по переданному имени
func (br *ByteRepo) Delete(name, userId string) (bool, error) {
	err := config.DB.
		Table("bytes as b").
		Where("b.name = ? and b.user_id = ?", name, userId).
		Delete(&domain.Byte{}).Error

	if err != nil {
		return false, err
	}

	return true, err
}

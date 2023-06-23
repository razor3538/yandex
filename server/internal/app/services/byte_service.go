package services

import (
	"github.com/gofrs/uuid"
	repositories "server/internal/app/repository"
	"server/internal/domain"
	"server/internal/models"
	"server/internal/tools"
)

// ByteService структура
type ByteService struct{}

// NewByteService метод возвращает указатель на структуру ByteService со всеми ее методами
func NewByteService() *ByteService {
	return &ByteService{}
}

var byteRepo = repositories.NewByteRepo()

// Save сохраняет текст
func (bs *ByteService) Save(byteModel models.SaveBytes, userId string) (domain.Byte, error) {
	id, err := uuid.FromString(userId)

	var bytes = domain.Byte{
		Base:   domain.Base{},
		Bytes:  byteModel.Bytes,
		Meta:   byteModel.Meta,
		Name:   byteModel.Name,
		UserId: id,
	}

	bytes.Bytes = string(tools.Base64Encode([]byte(bytes.Bytes)))

	result, err := byteRepo.Save(bytes)

	if err != nil {
		return domain.Byte{}, err
	}

	return result, nil
}

// GetById возвращает текст по переданному имени
func (bs *ByteService) GetById(name string, userId string) (domain.Byte, error) {
	result, err := byteRepo.GetByKey(name, userId)
	if err != nil {
		return domain.Byte{}, err
	}

	return result, nil
}

// Delete удаляет текст по переданному имени
func (bs *ByteService) Delete(name string, userId string) (bool, error) {
	result, _ := byteRepo.Delete(name, userId)

	return result, nil
}

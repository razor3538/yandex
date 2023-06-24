package api

import (
	"errors"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	middleware "server/internal/app/midleware"
	"server/internal/app/services"
	"server/internal/models"
	"server/internal/tools"
)

// BytePI структура
type BytePI struct{}

// NewByteAPI метод возвращает указатель на структуру BytePI со всеми ее методами
func NewByteAPI() *BytePI {
	return &BytePI{}
}

var byteService = services.NewByteService()

// Save сохраняет текст
func (ba *BytePI) Save(c *gin.Context) {
	var body models.SaveBytes

	headerToken := c.GetHeader("Authorization")

	if headerToken == "" {
		tools.CreateError(http.StatusNoContent, errors.New("пустой токен"), c)
		return
	}

	value, _ := middleware.Passport().ParseTokenString(headerToken)

	id := jwt.ExtractClaimsFromToken(value)["id"]

	if err := tools.RequestBinderBody(&body, c); err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	byteModel, err := byteService.Save(body, id.(string))

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusCreated, byteModel)
}

// Get возвращает текст по переданному имени
func (ba *BytePI) Get(c *gin.Context) {
	var body models.GetBytes

	headerToken := c.GetHeader("Authorization")

	if headerToken == "" {
		tools.CreateError(http.StatusNoContent, errors.New("пустой токен"), c)
		return
	}

	value, _ := middleware.Passport().ParseTokenString(headerToken)

	id := jwt.ExtractClaimsFromToken(value)["id"]

	if err := tools.RequestBinderBody(&body, c); err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	byteModel, err := byteService.GetById(body.Name, id.(string))

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	textRes, err := tools.Base64Decode([]byte(byteModel.Bytes))
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("%s: %s", byteModel.Name, textRes))
}

// Delete удаляет текст по переданному имени
func (ba *BytePI) Delete(c *gin.Context) {
	var body models.GetBytes

	headerToken := c.GetHeader("Authorization")

	if headerToken == "" {
		tools.CreateError(http.StatusNoContent, errors.New("пустой токен"), c)
		return
	}

	value, _ := middleware.Passport().ParseTokenString(headerToken)

	id := jwt.ExtractClaimsFromToken(value)["id"]

	if err := tools.RequestBinderBody(&body, c); err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	_, err := byteService.Delete(body.Name, id.(string))

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusOK, "Удаление прошло успешно")
}

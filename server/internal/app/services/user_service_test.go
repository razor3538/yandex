package services

import (
	"server/config"
	"server/internal/domain"
	"server/internal/models"
	"testing"
)

func TestUserService_Save(t *testing.T) {
	config.InitEnv()
	config.InitDB()

	tests := []struct {
		name    string
		model   models.SaveUserRequest
		want    domain.User
		wantErr bool
	}{
		{
			name: "valid test",
			model: models.SaveUserRequest{
				Login:    "admin@mail.ru",
				Password: "123321",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &UserService{}
			model, err := us.Save(tt.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if model.Login != model.Login {
				t.Errorf("admin@mail.ru not the same with %v", model.Login)
				return
			}
		})
	}
}

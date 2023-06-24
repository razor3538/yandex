package services

import (
	"reflect"
	"server/config"
	"server/internal/domain"
	"server/internal/models"
	"server/internal/tools"
	"testing"
)

func TestPasswordService_Save(t *testing.T) {
	config.InitDB()

	type args struct {
		passModel models.SavePassword
		userId    string
	}
	tests := []struct {
		name    string
		args    args
		want    domain.Password
		wantErr bool
	}{
		{
			name: "valid test",
			args: args{
				passModel: models.SavePassword{
					Login:    "admin",
					Password: "123",
					Meta:     "",
					Name:     "admin login",
				},
				userId: "e0048afb-bcb0-4ea7-ac30-a9eb04815b4d",
			},
			wantErr: false,
			want: domain.Password{
				Name:     "admin login",
				Login:    string(tools.Base64Encode([]byte("admin"))),
				Password: string(tools.Base64Encode([]byte("123"))),
				Meta:     "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &PasswordService{}
			got, err := ps.Save(tt.args.passModel, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Login, tt.want.Login) {
				t.Errorf("Save() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPasswordService_GetById(t *testing.T) {
	config.InitDB()

	type args struct {
		name   string
		userId string
	}
	tests := []struct {
		name    string
		args    args
		want    domain.Password
		wantErr bool
	}{
		{
			name: "valid test",
			args: args{
				name:   "admin login",
				userId: "e0048afb-bcb0-4ea7-ac30-a9eb04815b4d",
			},
			wantErr: false,
			want: domain.Password{
				Name:     "admin login",
				Login:    string(tools.Base64Encode([]byte("admin"))),
				Password: string(tools.Base64Encode([]byte("123"))),
				Meta:     "",
			},
		},

		{
			name: "invalid test",
			args: args{
				name:   "hellos",
				userId: "e0048afb-bcb0-4ea7-ac30-a9eb04815b4d",
			},
			wantErr: true,
		},

		{
			name: "invalid test",
			args: args{
				name:   "admin login",
				userId: "q0048afb-bcb0-4ea7-ac30-a9eb04815b4d",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &PasswordService{}
			got, err := ps.GetById(tt.args.name, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Login, tt.want.Login) {
				t.Errorf("GetById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPasswordService_Delete(t *testing.T) {
	config.InitDB()

	type args struct {
		name   string
		userId string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "invalid test",
			args: args{
				name:   "hellos",
				userId: "e0048afb-bcb0-4ea7-ac30-a9eb04815b4d",
			},
			wantErr: false,
			want:    true,
		},

		{
			name: "invalid test",
			args: args{
				name:   "admino login",
				userId: "w0048afb-bcb0-4ea7-ac30-a9eb04815b4d",
			},
			wantErr: false,
			want:    true,
		},
		{
			name: "valid test",
			args: args{
				name:   "admin login",
				userId: "e0048afb-bcb0-4ea7-ac30-a9eb04815b4d",
			},
			wantErr: false,
			want:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &PasswordService{}
			got, err := ps.Delete(tt.args.name, tt.args.userId)

			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Delete() got = %v, want %v", got, tt.want)
			}
		})
	}
}

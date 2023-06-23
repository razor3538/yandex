package services

import (
	"reflect"
	"server/config"
	"server/internal/domain"
	"server/internal/models"
	"server/internal/tools"
	"testing"
)

func TestCardService_Save(t *testing.T) {
	config.InitDB()

	type args struct {
		cardModel models.SaveCard
		userId    string
	}
	tests := []struct {
		name    string
		args    args
		want    domain.Cards
		wantErr bool
	}{
		{
			name: "valid test",
			args: args{
				cardModel: models.SaveCard{
					Number:  "1111 1111 1111 1111",
					DateEnd: "24/11",
					CVS:     "111",
					Bank:    "СберБанк",
					Meta:    "",
					Name:    "карта сбера",
				},
				userId: "e0048afb-bcb0-4ea7-ac30-a9eb04815b4d",
			},
			wantErr: false,
			want: domain.Cards{
				Name:    "карта сбера",
				Number:  string(tools.Base64Encode([]byte("1111 1111 1111 1111"))),
				DateEnd: string(tools.Base64Encode([]byte("24/11"))),
				CVS:     string(tools.Base64Encode([]byte("111"))),
				Bank:    string(tools.Base64Encode([]byte("СберБанк"))),
				Meta:    "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &CardService{}
			got, err := cs.Save(tt.args.cardModel, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Number, tt.want.Number) {
				t.Errorf("Save() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardService_GetById(t *testing.T) {
	config.InitDB()

	type args struct {
		name   string
		userId string
	}
	tests := []struct {
		name    string
		args    args
		want    domain.Cards
		wantErr bool
	}{
		{
			name: "valid test",
			args: args{
				name:   "карта сбера",
				userId: "e0048afb-bcb0-4ea7-ac30-a9eb04815b4d",
			},
			wantErr: false,
			want: domain.Cards{
				Name:    "карта сбера",
				Number:  string(tools.Base64Encode([]byte("1111 1111 1111 1111"))),
				DateEnd: string(tools.Base64Encode([]byte("24/11"))),
				CVS:     string(tools.Base64Encode([]byte("111"))),
				Bank:    string(tools.Base64Encode([]byte("СберБанк"))),
				Meta:    "",
			},
		},

		{
			name: "invalid test",
			args: args{
				name:   "карта сбербанка",
				userId: "e0048afb-bcb0-4ea7-ac30-a9eb04815b4d",
			},
			wantErr: true,
		},

		{
			name: "invalid test",
			args: args{
				name:   "карта сбера",
				userId: "q0048afb-bcb0-4ea7-ac30-a9eb04815b4d",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &CardService{}
			got, err := cs.GetById(tt.args.name, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Number, tt.want.Number) {
				t.Errorf("GetById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardService_Delete(t *testing.T) {
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
				name:   "карта сбер",
				userId: "e0048afb-bcb0-4ea7-ac30-a9eb04815b4d",
			},
			wantErr: false,
			want:    true,
		},

		{
			name: "invalid test",
			args: args{
				name:   "карта сб",
				userId: "w0048afb-bcb0-4ea7-ac30-a9eb04815b4d",
			},
			wantErr: false,
			want:    true,
		},
		{
			name: "valid test",
			args: args{
				name:   "карта сбера",
				userId: "e0048afb-bcb0-4ea7-ac30-a9eb04815b4d",
			},
			wantErr: false,
			want:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &CardService{}
			got, err := cs.Delete(tt.args.name, tt.args.userId)
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

package services

import (
	"reflect"
	"server/config"
	"server/internal/domain"
	"server/internal/models"
	"server/internal/tools"
	"testing"
)

func TestTextService_Save(t *testing.T) {
	config.InitDB()

	type args struct {
		textModel models.SaveText
		userId    string
	}
	tests := []struct {
		name    string
		args    args
		want    domain.Text
		wantErr bool
	}{
		{
			name: "valid test",
			args: args{
				textModel: models.SaveText{
					Text: "hello world",
					Meta: "",
					Name: "hello text",
				},
				userId: "e0048afb-bcb0-4ea7-ac30-a9eb04815b4d",
			},
			wantErr: false,
			want: domain.Text{
				Name: "hello",
				Text: string(tools.Base64Encode([]byte("hello world"))),
				Meta: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &TextService{}
			got, err := ts.Save(tt.args.textModel, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Text, tt.want.Text) {
				t.Errorf("Save() got = %v, want %v", got.Text, tt.want.Text)
			}
		})
	}
}

func TestTextService_GetById(t *testing.T) {
	config.InitDB()

	type args struct {
		name   string
		userId string
	}
	tests := []struct {
		name    string
		args    args
		want    domain.Text
		wantErr bool
	}{
		{
			name: "valid test",
			args: args{
				name:   "hello text",
				userId: "e0048afb-bcb0-4ea7-ac30-a9eb04815b4d",
			},
			wantErr: false,
			want: domain.Text{
				Name: "hello",
				Text: string(tools.Base64Encode([]byte("hello world"))),
				Meta: "",
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
				name:   "hello text",
				userId: "q0048afb-bcb0-4ea7-ac30-a9eb04815b4d",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &TextService{}
			got, err := ts.GetById(tt.args.name, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Text, tt.want.Text) {
				t.Errorf("GetById() got = %v, want %v", got.Text, tt.want.Text)
				return
			}
		})
	}
}

func TestTextService_Delete(t *testing.T) {
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
				name:   "hellop",
				userId: "w0048afb-bcb0-4ea7-ac30-a9eb04815b4d",
			},
			wantErr: false,
			want:    true,
		},
		{
			name: "valid test",
			args: args{
				name:   "hello text",
				userId: "e0048afb-bcb0-4ea7-ac30-a9eb04815b4d",
			},
			wantErr: false,
			want:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &TextService{}
			got, err := ts.Delete(tt.args.name, tt.args.userId)
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

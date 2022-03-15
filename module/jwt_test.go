package module

import (
	"testing"
)

func TestGenreateToken(t *testing.T) {
	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "case 1",
			args:    args{userID: 1},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenreateToken(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == tt.want {
				t.Errorf("GenreateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerifyToken(t *testing.T) {
	s, _ := GenreateToken(1)
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "case 2",
			args:    args{token: s},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VerifyToken(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.UserID != tt.want {
				t.Errorf("VerifyToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

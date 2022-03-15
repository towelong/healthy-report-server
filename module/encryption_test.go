package module

import (
	"testing"
)

func TestNewEncryption(t *testing.T) {
	type args struct {
		password string
		size     int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Case1",
			args: args{
				password: "123123",
				size:     3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEncryption(WithPassword(tt.args.password)); got.EncodePassword() == "" {
				t.Errorf("NewEncryption() = %v, want %v", got, tt.want)
			}
			if got := NewEncryption(WithPassword(tt.args.password)); got.VerifyPassword(got.Password) == tt.want {
				t.Errorf("NewEncryption() = %v, want %v", got, tt.want)
			}
		})
	}
}

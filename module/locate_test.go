package module

import (
	"testing"
)

func Test_getLocationByAddress(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    *Location
		wantErr bool
	}{
		{
			name: "Case 1",
			args: args{address: "江西省九江市共青城市江西农业大学南昌商学院"},
			want: &Location{
				Status: "OK",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getLocationByAddress(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("getLocationByAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want.Status != got.Status {
				t.Errorf("getLocationByAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

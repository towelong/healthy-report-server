package module

import (
	"testing"
)

func Test_splitAddress(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
		want *AddressDetail
	}{
		{
			name: "Case 1",
			args: args{address: "江西省九江市共青城市青年大道79号江西农业大学南昌商学院"},
			want: &AddressDetail{
				Province: "江西省",
				City:     "九江市",
				District: "共青城市",
				Street:   "青年大道79号江西农业大学南昌商学院",
			},
		},
		{
			name: "Case 2",
			args: args{address: "江西省赣州市宁都县宁都第三小学"},
			want: &AddressDetail{
				Province: "江西省",
				City:     "赣州市",
				District: "宁都县",
				Street:   "宁都第三小学",
			},
		},
		{
			name: "Case 3",
			args: args{address: "江西省赣州市章贡区黄金大道99号"},
			want: &AddressDetail{
				Province: "江西省",
				City:     "赣州市",
				District: "章贡区",
				Street:   "黄金大道99号",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitAddress(tt.args.address); got.Province != tt.want.Province {
				t.Errorf("Province = %v, want %v", got.Province, tt.want.Province)
			}
			if got := splitAddress(tt.args.address); got.City != tt.want.City {
				t.Errorf("City = %v, want %v", got.City, tt.want.City)
			}
			if got := splitAddress(tt.args.address); got.District != tt.want.District {
				t.Errorf("District = %v, want %v", got.District, tt.want.District)
			}
		})
	}
}

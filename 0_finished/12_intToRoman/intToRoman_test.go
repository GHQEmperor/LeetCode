package inttoroman

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				num: 1994,
			},
			want: "MICMIXCII",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}

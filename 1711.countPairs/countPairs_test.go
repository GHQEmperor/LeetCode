package countpairs

import "testing"

func Test_countPairs(t *testing.T) {
	type args struct {
		deliciousness []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				deliciousness: []int{1, 3, 5, 7, 9},
			},
			want: 4,
		},
		{
			name: "2",
			args: args{
				deliciousness: []int{1, 1, 1, 3, 3, 3, 7},
			},
			want: 15,
		},
		{
			name: "3",
			args: args{
				deliciousness: []int{2160, 1936, 3, 29, 27, 5, 2503, 1593, 2, 0, 16, 0, 3860, 28908, 6, 2, 15, 49, 6246, 1946, 23, 105, 7996, 196, 0, 2, 55, 457, 5, 3, 924, 7268, 16, 48, 4, 0, 12, 116, 2628, 1468},
			},
			want: 53,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPairs(tt.args.deliciousness); got != tt.want {
				t.Errorf("countPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

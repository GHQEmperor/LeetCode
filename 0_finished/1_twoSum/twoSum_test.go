package twosum

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nil",
			args: args{
				nums:   []int{},
				target: 1,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

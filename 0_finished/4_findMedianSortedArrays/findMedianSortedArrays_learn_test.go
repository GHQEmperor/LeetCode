package findmediansortedarrays

import "testing"

func Test_findMedianSortedArraysLearn(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "1",
			args: args{
				nums1: []int{1, 2},
				nums2: []int{3, 4},
			},
			want: 2.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArraysLearn(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArraysLearn() = %v, want %v", got, tt.want)
			}
		})
	}
}

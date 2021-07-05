package maxlevelsum

import "testing"

func Test_maxLevelSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1,7,0,7,-8,null,null]",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{Val: 7,
					Left:  &TreeNode{Val: 7},
					Right: &TreeNode{Val: -8},
				},
				Right: &TreeNode{Val: 0}}},
			want: 2,
		},
		{
			name: "[989,null,10250,98693,-89388,null,null,null,-32127]",
			args: args{root: &TreeNode{
				Val:  989,
				Left: nil,
				Right: &TreeNode{Val: 10250,
					Left:  &TreeNode{Val: 98693},
					Right: &TreeNode{Val: -89388, Right: &TreeNode{Val: -32127}},
				},
			}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxLevelSum(tt.args.root); got != tt.want {
				t.Errorf("maxLevelSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

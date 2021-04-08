package leetcode_test

// 题序：	153.寻找旋转排序数组中的最小值
// 题目链接：	https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/

import (
	"testing"
)

type Test struct {
	Input  []int
	Output int
}

func TestLeetCode(t *testing.T) {
	tests := []Test{
		{Input: []int{3, 4, 5, 1, 2}, Output: 1},
		{Input: []int{4, 5, 6, 7, 0, 1, 2}, Output: 0},
		{Input: []int{11, 13, 15, 17}, Output: 11},
	}

	for _, test := range tests {
		out := findMin(test.Input)
		if out != test.Output {
			t.Errorf("test %v\t, out %v\n", test.Output, out)
		}
	}
}

func findMin(nums []int) int {
	return 0
}

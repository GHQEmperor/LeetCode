package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 1}))
}

/*
	https://leetcode-cn.com/problems/container-with-most-water/
*/

func maxArea(height []int) int {
	hLen := len(height)
	var maxarea int
	left, right := 0, hLen-1
	for left < right {
		maxarea = max(min(height[left], height[right])*(right-left), maxarea)
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return maxarea
}

func min(v1, v2 int) int {
	if v1 > v2 {
		return v2
	}
	return v1
}

func max(v1, v2 int) int {
	if v1 < v2 {
		return v2
	}
	return v1
}

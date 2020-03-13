package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 2, 0, 4, 3, 1}
	fmt.Println(nums)
	nextPermutation(nums)

	fmt.Println(nums)
}

/*
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
必须原地修改，只允许使用额外常数空间。
以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
*/
func nextPermutation(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	var topI int
	for i := n - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			topI = i
			break
		}
	}
	//fmt.Println("top", topI)

	if topI == 0 {
		sort.Ints(nums)
		return
	}

	for i := n - 1; i >= topI; i-- {
		if nums[i] > nums[topI-1] {
			nums[i], nums[topI-1] = nums[topI-1], nums[i]
			k := len(nums) - 1
			for topI < k {
				nums[topI], nums[k] = nums[k], nums[topI]
				topI++
				k--
			}
			//sort.Ints(nums[topI:])
			break
		}
	}
}

package foursum

import (
	"sort"
)

/*
18. 四数之和
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，
使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
注意：
答案中不可以包含重复的四元组。

示例：
给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
*/

func fourSum(nums []int, target int) [][]int {
	n := len(nums)

	sort.Ints(nums)

	var result [][]int
	for i := 0; i < n-3; i++ {
		for j := n - 1; j > i+2; j-- {
			l, r := i+1, j-1
			for l < r {
				min := nums[i] + nums[i+1] + nums[i+2] + nums[i+3]
				if min > target {
					break
				}
				max := nums[j] + nums[j-1] + nums[j-2] + nums[j-3]
				if max < target {
					break
				}
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}

					l++
					r--
				} else if sum > target {
					r--
				} else {
					l++
				}
			}

			// ?
			for j > i+2 && nums[j] == nums[j-1] {
				j--
			}
		}

		for i < n-3 && nums[i] == nums[i+1] {
			i++
		}
	}
	return result

}

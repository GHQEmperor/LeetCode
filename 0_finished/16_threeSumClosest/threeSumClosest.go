package threesumclosest

import "sort"

/*
16. 最接近的三数之和
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，
使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
*/

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}

	sort.Ints(nums)
	goodSum := nums[0] + nums[1] + nums[2]

	for i := 0; i < n-2; i++ {
		l, r := i+1, n-1
		for l < r {
			min := nums[i] + nums[l] + nums[l+1]
			if min > target {
				if posit(target-goodSum) > posit(target-min) {
					goodSum = min
				}
				break
			}

			max := nums[i] + nums[r] + nums[r-1]
			if target > max {
				if posit(target-goodSum) > posit(target-max) {
					goodSum = max
				}
				break
			}

			sum := nums[i] + nums[l] + nums[r]
			if posit(target-sum) < posit(target-goodSum) {
				goodSum = sum
			}
			if sum == target {
				return sum
			}
			if sum > target {
				for l < r && nums[r] == nums[r+1] {
					r--
				}
				r--
			} else {
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				l++
			}
		}
		for i < n-2 && nums[i] == nums[i+1] {
			i++
		}
	}
	return goodSum
}

func posit(v int) int {
	if v < 0 {
		return -1 * v
	}
	return v
}

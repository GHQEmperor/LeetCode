package main

import "fmt"

/*
64. 最小路径和
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。
示例:
输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。
*/

func main() {
	test := [][][]int{
		{
			{1, 3, 1},
			{1, 5, 1},
			{4, 2, 1},
		},
	}
	for _, v := range test {
		fmt.Println(minPathSum(v))
	}
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	if m == 1 || n == 1 {
		var sum int
		for i := range grid {
			for j := range grid[i] {
				sum += grid[i][j]
			}
		}
		return sum
	}

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	print2(dp)

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

func min(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

func print2(v [][]int) {
	for _, row := range v {
		for _, value := range row {
			fmt.Printf("%d\t", value)
			//if value {
			//	fmt.Printf("√\t")
			//} else {
			//	fmt.Printf("×\t")
			//}
		}
		fmt.Println()
	}
	fmt.Println()
}

package maximalrectangle

/*
85. 最大矩形
给定一个仅包含 0 和 1 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。

示例:
输入:
[
  ["1","0","1","0","0"],
  ["1","0","1","1","1"],
  ["1","1","1","1","1"],
  ["1","0","0","1","0"]
]
输出: 6
*/

func maximalRectangle(matrix [][]byte) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	m := len(matrix[0])
	if m == 0 {
		return 0
	}
	if n == 1 {
		var max int
		for i := 0; i < m; i++ {
			for j := i; j < m; j++ {
				if matrix[0][j] == '0' {
					break
				}
				max = getMax(max, j-i+1)
			}
		}
		return max
	}
	if m == 1 {
		var max int
		for i := 0; i < n; i++ {
			for j := i; j < n; j++ {
				if matrix[j][0] == '0' {
					break
				}
				max = getMax(max, j-i+1)
			}
		}
		return max
	}

	l2, l1 := make([]int, m), make([]int, m)
	var max int
	for i := 1; i <= n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i-1][j] != '0' {
				l1[j] = l2[j] + 1
			} else {
				l1[j] = 0
			}
		}

		var lineMax int
		for k := 0; k < m; k++ {
			height := l1[k]
			for l := k; l < m; l++ {
				if l1[l] == 0 {
					break
				}
				height = getMin(height, l1[l])
				lineMax = getMax(lineMax, (l-k+1)*height)
			}
		}
		max = getMax(max, lineMax)
		l2, l1 = l1, l2
	}

	return max
}

func getMin(v1, v2 int) int {
	if v1 > v2 {
		return v2
	}
	return v1
}

func getMax(v1, v2 int) int {
	if v1 < v2 {
		return v2
	}
	return v1
}

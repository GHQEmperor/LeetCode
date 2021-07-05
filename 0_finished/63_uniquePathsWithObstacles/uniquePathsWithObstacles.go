package uniquePathsWithObstacles

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}
	if m == 1 || n == 1 {
		for i := range obstacleGrid {
			for j := range obstacleGrid[i] {
				if obstacleGrid[i][j] == 1 {
					return 0
				}
			}
		}
		return 1
	}
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 1; i < m; i++ {
		if (dp[i-1][0] == 1 || i == 1) && obstacleGrid[i][0] == 0 {
			dp[i][0] = 1
		} else {
			dp[i][0] = 0
		}
	}
	for i := 1; i < n; i++ {
		if (dp[0][i-1] == 1 || i == 1) && obstacleGrid[0][i] == 0 {
			dp[0][i] = 1
		} else {
			dp[0][i] = 0
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			} else {
				dp[i][j] = 0
			}
		}
	}
	return dp[m-1][n-1]
}

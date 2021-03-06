package mindistance

/*
72. 编辑距离
给定两个单词 word1 和 word2，计算出将 word1 转换成 word2 所使用的最少操作数 。
你可以对一个单词进行如下三种操作：
插入一个字符
删除一个字符
替换一个字符
示例 1:
输入: word1 = "horse", word2 = "ros"
输出: 3
解释:
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2:
输入: word1 = "intention", word2 = "execution"
输出: 5
解释:
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')
*/

func minDistance(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)
	if m < n {
		n, m = m, n
		word1, word2 = word2, word1
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
		dp[i][0] = i
	}
	for i := 1; i <= m; i++ {
		dp[0][i] = i
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]-1)
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			}
		}
	}
	return dp[n][m]
}

func min(v1, v2, v3 int) int {
	v := v1
	if v > v2 {
		v = v2
	}
	if v > v3 {
		v = v3
	}
	return v
}

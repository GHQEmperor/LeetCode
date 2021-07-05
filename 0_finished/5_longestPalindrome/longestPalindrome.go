package longestPalindrome

import (
	"fmt"
)

// todo: Manacher（马拉车）算法
// 	https://blog.csdn.net/daidaineteasy/article/details/86238047
/*
5. 最长回文子串
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"
*/

/*
当 i == j，dp[i][j] 是回文子串（单字符都是回文子串）；
当 j - i < 3，只要 S[i] == S[j]，则 dp[i][j] 是回文子串（如"aa"，“aba”），否则不是；
当 j - i >= 3，如果 S[i] == S[j] && dp[i+1][j-1] ，则 dp[i][j] 是回文子串，否则不是 。
*/

func main() {
	fmt.Println(longestPalindrome("bdcabbacdf"))
}

func longestPalindrome(s string) string {
	//return function.Dynamic(s)
	return dynamic(s)
	//return function.CenterExtension(s)
}

func dynamic(s string) string {
	sLen := len(s)

	if s == "" || sLen == 1 {
		return s
	}

	dp := make([][]bool, sLen)
	for i := 0; i < sLen; i++ {
		dp[i] = make([]bool, sLen)
	}

	/*
		当 i == j，dp[i][j] 是回文子串（单字符都是回文子串）；
		当 j - i < 3，只要 S[i] == S[j]，则 dp[i][j] 是回文子串（如"aa"，“aba”），否则不是；
		当 j - i >= 3，如果 S[i] == S[j] && dp[i+1][j-1] ，则 dp[i][j] 是回文子串，否则不是 。
	*/
	var left, right int
	for i := sLen - 1; i >= 0; i-- {
		dp[i][i] = true
		for j := i + 1; j < sLen; j++ {
			dp[i][j] = s[i] == s[j] && (j-i < 3 || dp[i+1][j-1])
			if dp[i][j] && right-left < j-i {
				left = i
				right = j
			}
		}
	}

	return s[left : right+1]
}

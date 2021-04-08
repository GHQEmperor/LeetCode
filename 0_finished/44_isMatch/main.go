package main

import "fmt"

type Param struct {
	s, p   string
	result bool
}

/*
44. 通配符匹配
给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。
'?' 可以匹配任何单个字符。
'*' 可以匹配任意字符串（包括空字符串）。
两个字符串完全匹配才算匹配成功。

说明:
s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。
示例 1:
输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
示例 2:
输入:
s = "aa"
p = "*"
输出: true
解释: '*' 可以匹配任意字符串。
示例 3:
输入:
s = "cb"
p = "?a"
输出: false
解释: '?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。
示例 4:
输入:
s = "adceb"
p = "*a*b"
输出: true
解释: 第一个 '*' 可以匹配空字符串, 第二个 '*' 可以匹配字符串 "dce".
示例 5:
输入:
s = "acdcb"
p = "a*c?b"
输入: false
*/

func main() {
	test := []Param{
		{"aa", "a", false},
		//{"", "a", false},
		{"aabbc", "*ab*c", true},
	}
	for _, v := range test {
		fmt.Println(isMatch(v.s, v.p), v.result)
	}
}

func isMatch(s string, p string) bool {
	ns := len(s)
	np := len(p)
	dp := make([][]bool, np+1)
	for i := range dp {
		dp[i] = make([]bool, ns+1)
	}
	dp[0][0] = true
	for i := 1; i <= ns; i++ {
		dp[0][i] = false
	}
	for i := 1; i <= np; i++ {
		if p[i-1] == '*' && dp[i-1][0] {
			dp[i][0] = true
		}
	}
	print2(dp)

	for i := 1; i <= np; i++ {
		for j := 1; j <= ns; j++ {
			if ((p[i-1] == s[j-1] || p[i-1] == '?' || p[i-1] == '*') && dp[i-1][j-1]) ||
				(p[i-1] == '*' && (dp[i-1][j] || dp[i][j-1])) {
				dp[i][j] = true
			}
		}
	}
	print2(dp)
	fmt.Println("np:", np+1, "ns:", ns+1)
	return dp[np][ns]
}

func print2(v [][]bool) {
	for _, row := range v {
		for _, value := range row {
			if value {
				fmt.Printf("√\t")
			} else {
				fmt.Printf("×\t")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
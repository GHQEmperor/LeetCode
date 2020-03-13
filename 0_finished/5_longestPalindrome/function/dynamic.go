package function

import "fmt"

func Dynamic(s string) string {
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
		//fmt.Printf("%#v\n", dp)
		printf(s, dp)
	}

	return s[left : right+1]
}

func printf(s string, dp [][]bool) {
	fmt.Printf("\ti\t")
	for i := range dp {
		fmt.Printf("%c\t", s[i])
	}
	fmt.Println()
	fmt.Printf("j\t\t")
	for i := range dp {
		fmt.Printf("%d\t", i)
	}
	fmt.Println()
	for j, v := range dp {
		//for j := range v {
		fmt.Printf("%c\t", s[j])
		fmt.Printf("%d\t", j)
		//}
		for _, v2 := range v {
			if v2 {
				fmt.Printf("√\t")
			} else {
				fmt.Printf("×\t")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}

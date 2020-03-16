package main

import (
	"LeetCode/tools"
	"fmt"
)

/*
87. 扰乱字符串
给定一个字符串 s1，我们可以把它递归地分割成两个非空子字符串，从而将其表示为二叉树。
下图是字符串 s1 = "great" 的一种可能的表示形式。
    great
   /    \
  gr    eat
 / \    /  \
g   r  e   at
           / \
          a   t
在扰乱这个字符串的过程中，我们可以挑选任何一个非叶节点，然后交换它的两个子节点。
例如，如果我们挑选非叶节点 "gr" ，交换它的两个子节点，将会产生扰乱字符串 "rgeat" 。
    rgeat
   /    \
  rg    eat
 / \    /  \
r   g  e   at
           / \
          a   t
我们将 "rgeat” 称作 "great" 的一个扰乱字符串。
同样地，如果我们继续交换节点 "eat" 和 "at" 的子节点，将会产生另一个新的扰乱字符串 "rgtae" 。
    rgtae
   /    \
  rg    tae
 / \    /  \
r   g  ta  e
       / \
      t   a
我们将 "rgtae” 称作 "great" 的一个扰乱字符串。
给出两个长度相等的字符串 s1 和 s2，判断 s2 是否是 s1 的扰乱字符串。
示例 1:
输入: s1 = "great", s2 = "rgeat"
输出: true
示例 2:
输入: s1 = "abcde", s2 = "caebd"
输出: false
*/

func main() {
	test := [][]interface{}{
		{"great", "rgeat", true},
		{"abcde", "caebd", false},
	}
	for _, v := range test {
		fmt.Println(isScramble(v[0].(string), v[1].(string)), v[2].(bool))
	}
}

func isScramble(s1 string, s2 string) bool {
	n := len(s1)
	//hashMap := make(map[string]struct{})
	for i := 0; i < n; i++ {
		res := get([]byte(s1[0:i]), []byte(s1[i:]))
		for _, v := range res {
			tools.PrintSlice(v)
			if string(v) == s2 {
				return true
			}
		}
	}
	return false
}

func get(header, tail []byte) [][]byte {
	n := len(header)
	m := len(tail)
	var headers, tails [][]byte
	if n > 2 {
		for i := 0; i < n; i++ {
			headers = append(headers, get(header[0:i], header[i:])...)
			headers = append(headers, get(header[i:], header[0:i])...)
		}
	}
	if m > 2 {
		for i := 0; i < n; i++ {
			tails = append(tails, get(tail[0:i], tail[i:])...)
			tails = append(tails, get(tail[i:], tail[0:i])...)
		}
	}
	nn := len(headers)
	if nn == 0 || headers == nil {
		return tails
	}
	mm := len(tails)
	if mm == 0 || tails == nil {
		return headers
	}
	res := make([][]byte, 0, nn*mm)
	for i := 0; i < nn; i++ {
		for j := 0; j < mm; j++ {
			add := append([]byte{}, headers[i]...)
			add = append(add, tails[j]...)
			res = append(res, add)
		}
	}
	return res
}

package main

import "fmt"

func main() {
	//a := []int{1, 2, 3, 4, 5}
	//fmt.Println(a[:4])
	fmt.Println(isValid("()"))
}

/*
20. 有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:
输入: "()"
输出: true

示例 2:
输入: "()[]{}"
输出: true

示例 3:
输入: "(]"
输出: false

示例 4:
输入: "([)]"
输出: false

示例 5:
输入: "{[]}"
输出: true
*/
func isValid(s string) bool {
	n := len(s)
	stack := make([]byte, 0, n)
	for _, v := range s {
		if v == '{' || v == '[' || v == '(' {
			stack = add(stack, byte(v))
			continue
		}
		switch v {
		case '}':
			if getTop(stack) != '{' {
				return false
			}
			stack = delTop(stack)
		case ']':
			if getTop(stack) != '[' {
				return false
			}
			stack = delTop(stack)
		case ')':
			if getTop(stack) != '(' {
				return false
			}
			stack = delTop(stack)
		default:
			return false
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func delTop(l []byte) []byte {
	n := len(l)
	if n != 0 {
		return l[:n-1]
	}
	return nil
}

func getTop(l []byte) byte {
	n := len(l)
	if n == 0 {
		return 0
	}
	return l[n-1]
}

func add(l []byte, v byte) []byte {
	return append(l, v)
}

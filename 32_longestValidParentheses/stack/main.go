package main

import "fmt"

func main() {
	fmt.Println(longestValidParentheses(")()())"))
}

/*
32. 最长有效括号
给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

示例 1:
输入: "(()"
输出: 2
解释: 最长有效括号子串为 "()"
示例 2:
输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"
*/
func longestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}
	n := len(s)
	var stack Stack
	for i := range s {
		v, _ := stack.Top()
		//if s[i] == '(' {
		stack.Push(i)
		//} else if s[i] == ')' {
		//v, _ := stack.Top()
		//if index == -1 {
		//	stack.index = 0
		//	stack.content = nil
		//}
		if s[i] == ')' && v != -1 && s[v] == '(' {
			fmt.Println(v, i)
			stack.Pop()
			stack.Pop()
		}
		//}
	}
	var max int
	top, _ := stack.Top()
	if top == -1 {
		return n
	}
	max = n - 1 - top
	sn := len(stack.content)
	for i := sn - 1; i > 0; i-- {
		fmt.Println(i)
		count := (stack.content[i] - stack.content[i-1]) / 2 * 2
		if count > max {
			max = count
		}
	}
	count := stack.content[0] / 2 * 2
	if count > max {
		max = count
	}
	fmt.Println(stack)

	return max
}

type Stack struct {
	content []int
	index   int
}

func (s *Stack) Pop() int {
	if s.index < 0 {
		return -1
	}
	v := s.content[s.index]
	if s.index != 0 {
		s.content = s.content[:s.index]
	} else {
		s.content = nil
	}
	s.index--
	return v
}

func (s *Stack) Top() (v int, index int) {
	if s.index < 0 || s.content == nil {
		return -1, -1
	}
	//fmt.Println(s.content, s.index)
	return s.content[s.index], s.index
}

func (s *Stack) Push(v int) {
	if s.content == nil {
		s.index = -1
	}
	s.content = append(s.content, v)
	s.index++
}

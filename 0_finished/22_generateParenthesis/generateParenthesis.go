package generateParenthesis

/*
22. 括号生成
给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
例如，给出 n = 3，生成结果为：
[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/
func generateParenthesis(n int) []string {
	return gen([]byte(""), n, n)
}

func gen(res []byte, left, right int) []string {
	if right < left {
		return nil
	}
	var result []string
	if left == 0 {
		for right != 0 {
			res = append(res, ')')
			right--
		}
		return []string{string(res)}
	}

	result = append(result, gen(append(res, '('), left-1, right)...)
	result = append(result, gen(append(res, ')'), left, right-1)...)
	return result
}

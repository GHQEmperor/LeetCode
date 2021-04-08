package main

import "fmt"

/*
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);
示例 1:

输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
示例 2:

输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G
*/

func main() {
	fmt.Printf("result:\n%s",
		convert("LEETCODEISHIRING", 3),
	)
}

// LCIRETOESIIGEDHN
// LECDIHRNETOESIIG
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	sLen := len(s)
	round := numRows*2 - 2
	width := round * 2
	result := make([][]byte, numRows)
	for i := range result {
		result[i] = make([]byte, 0, width)
	}

	var flag bool
	for i, j := 0, 0; i < sLen; i++ {
		result[j] = append(result[j], s[i])

		if !flag {
			j++
		} else {
			j--
		}

		if j >= numRows-1 || j <= 0 {
			flag = !flag
		}
	}

	resultBytes := make([]byte, 0, sLen)
	//var resultStr string
	for _, v := range result {
		//resultStr += string(v)
		resultBytes = append(resultBytes, v...)
	}

	return string(resultBytes)
}
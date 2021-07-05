package divide

/*
29. 两数相除
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
返回被除数 dividend 除以除数 divisor 得到的商。

示例 1:
输入: dividend = 10, divisor = 3
输出: 3

示例 2:
输入: dividend = 7, divisor = -3
输出: -2
*/
func divide(dividend int, divisor int) int {
	dividendFlag, divisorFlag := 1, 1
	if dividend < 0 {
		dividendFlag = -1
		dividend = 0 - dividend
	}
	if divisor < 0 {
		divisorFlag = -1
		divisor = 0 - divisor
	}
	var result int
	plus := 1
	div := divisor
	for dividend >= divisor {
		if dividend < div {
			div = divisor
			plus = 1
		}
		result += plus
		dividend -= div
		plus += plus
		div += div
	}

	if (dividendFlag == 1 && divisorFlag == 1) || (dividendFlag == -1 && divisorFlag == -1) {
		if result < 2147483648 {
			return result
		} else {
			return 2147483647
		}

	}

	result = 0 - result
	return result
}

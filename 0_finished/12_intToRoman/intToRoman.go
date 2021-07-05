package intToRoman

func intToRoman(num int) string {
	var result []byte

	for num != 0 {
		if num >= 1000 {
			result = append(result, 'M')
			num -= 1000
		} else if num >= 900 {
			result = append(result, 'C')
			result = append(result, 'M')
			num -= 900
		} else if num >= 500 {
			result = append(result, 'D')
			num -= 500
		} else if num >= 400 {
			result = append(result, 'C')
			result = append(result, 'D')
			num -= 400
		} else if num >= 100 {
			result = append(result, 'C')
			num -= 100
		} else if num >= 90 {
			result = append(result, 'X')
			result = append(result, 'C')
			num -= 90
		} else if num >= 50 {
			result = append(result, 'L')
			num -= 50
		} else if num >= 40 {
			result = append(result, 'X')
			result = append(result, 'L')
			num -= 40
		} else if num >= 10 {
			result = append(result, 'X')
			num -= 10
		} else if num >= 9 {
			result = append(result, 'I')
			result = append(result, 'X')
			num -= 9
		} else if num >= 5 {
			result = append(result, 'V')
			num -= 5
		} else if num >= 4 {
			result = append(result, 'I')
			result = append(result, 'V')
			num -= 4
		}
		result = append(result, 'I')
		num -= 1
	}
	return string(result)
}

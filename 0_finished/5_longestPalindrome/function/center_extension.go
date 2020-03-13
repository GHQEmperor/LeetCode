package function

func CenterExtension(s string) string {
	sLen := len(s)
	if s == "" || sLen == 1 {
		return s
	}

	var start, end int

	for i := 0; i < sLen-1; i++ {
		len1 := _func(s, i, i)

		len2 := _func(s, i, i+1)

		_len := max(len1, len2)

		if _len > end-start {
			start = i - (_len-1)/2
			end = i + _len/2
		}
	}

	return s[start : end+1]
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

func _func(s string, left, right int) int {
	sLen := len(s)
	for left > 0 && right < sLen && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

package longestcommonprefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	minLen := len(strs[0])
	for _, v := range strs {
		vLen := len(v)
		if vLen < minLen {
			minLen = vLen
		}
	}

	result := make([]byte, 0, minLen)
	for i := 0; i < minLen; i++ {
		value := strs[0][i]
		for _, v := range strs {
			if value != v[i] {
				return string(result)
			}
		}
		result = append(result, value)
	}
	return string(result)
}

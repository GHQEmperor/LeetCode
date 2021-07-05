package strStr

/*
实现strStr()函数。
给定一个haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回 -1。

示例 1:
输入: haystack = "hello", needle = "ll"
输出: 2

示例 2:
输入: haystack = "aaaaa", needle = "bba"
输出: -1
*/
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	if haystack == "" {
		return -1
	}
	nh := len(haystack)
	nn := len(needle)
	if nn > nh {
		return -1
	}

	for i := range haystack {
		var notAll bool
		for j := 0; j < nn; j++ {
			if i+j >= nh {
				return -1
			}
			if haystack[i+j] == needle[j] {
				continue
			} else {
				notAll = true
				break
			}

		}
		if !notAll {
			return i
		}
	}
	return -1
}

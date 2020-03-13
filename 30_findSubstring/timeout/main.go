package timeout

import "fmt"

func main() {
	fmt.Println(findSubstring("barfoofoobarthefoobarman", []string{"foo", "bar", "the"}))
}

func findSubstring(s string, words []string) []int {
	if s == "" || len(words) == 0 {
		return nil
	}
	var result []int
	bs := make(map[string]struct{})
	helper(words, 0, &bs)
	fmt.Println(bs)
	for key, _ := range bs {
		value := find(s, key)
		fmt.Println("key", key, "value", value)
		result = append(result, value...)
	}
	return result
}

func helper(bt []string, start int, strs *map[string]struct{}) {
	if start == len(bt) {
		var res string
		for _, v := range bt {
			res += v
		}
		(*strs)[res] = struct{}{}
	} else {
		for i := start; i <= len(bt)-1; i++ {
			if i != start {
				tmp := bt[start]
				bt[start] = bt[i]
				bt[i] = tmp
			}
			helper(bt, start+1, strs)

			if i != start {
				tmp := bt[start]
				bt[start] = bt[i]
				bt[i] = tmp
			}
		}
	}
}

// todo: 缺少最后一次匹配
func find(a, b string) []int {
	n, m := len(a), len(b)
	if n < m {
		return nil
	}

	var res []int
	for i := 0; i+m <= n; i++ {
		if a[i:i+m] == b {
			res = append(res, i)
		}
	}
	return res
}

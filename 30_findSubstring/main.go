package main

import "fmt"

func main() {
	fmt.Println(findSubstring("barfoofoobarthefoobarman", []string{"foo", "bar", "the"}))
}

// todo: 未完成
func findSubstring(s string, words []string) []int {
	if s == "" || len(words) == 0 {
		return nil
	}
	//n, m, mm := len(s), len(words), len(words[0])
	//wordsMap := make(map[string]int)
	//for _, v := range words {
	//	wordsMap[v]++
	//}

	//for i := 0; i < n; i++ {
	//getMap := make(map[string]int)
	//ls := s[i : i+m*mm+1]
	//for j := 0; j < i+m*mm+1; j+= {
	//	if
	//}
	//}

	return nil
}

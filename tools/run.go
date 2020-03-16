package tools

import "fmt"

func fn(v []int) int {
	return 0
}

func main() {
	test := [][]int{
		{-2, 1, -3, 4, -1, 2, 1, -5, 4},
	}
	for _, v := range test {
		fmt.Println(fn(v))
	}
}

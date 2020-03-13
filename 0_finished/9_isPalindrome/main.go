package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(12321))
}

func isPalindrome(x int) bool {
	return yes(x)
}

func yes(x int) bool {
	if x < 0 {
		return false
	}
	before := x
	var temp int
	for before != 0 {
		if temp != 0 {
			temp *= 10
		}
		temp += before % 10
		before /= 10
	}

	if temp == x {
		return true
	}
	return false
}

func no(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}

	var l []byte
	for x != 0 {
		l = append(l, byte(x%10))
		x /= 10
	}
	if l == nil {
		return false
	}
	//fmt.Println(l)
	lLen := len(l)
	lLenHalf := lLen / 2
	var count int
	for i := 0; i < lLenHalf; i++ {
		if l[i] == l[lLen-i-1] {
			count++
			continue
		}
	}
	if count == lLenHalf {
		return true
	}
	return false
}

package ispalindrome

func isPalindrome(x int) bool {
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

	return temp == x
}

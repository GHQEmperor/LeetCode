package romanToInt

func romanToInt(s string) int {
	bytes := []byte(s)
	bytesLen := len(bytes)
	var num int
	for i := 0; i < bytesLen; i++ {
		item := bytes[i]
		if item == 'M' {
			num += 1000
		} else if item == 'D' {
			num += 500
		} else if item == 'C' {
			if i+1 < bytesLen && bytes[i+1] == 'M' {
				i++
				num += 900
			} else if i+1 < bytesLen && bytes[i+1] == 'D' {
				i++
				num += 400
			} else {
				num += 100
			}
		} else if item == 'X' {
			if i+1 < bytesLen && bytes[i+1] == 'C' {
				i++
				num += 90
			} else if i+1 < bytesLen && bytes[i+1] == 'L' {
				i++
				num += 40
			} else {
				num += 10
			}
		} else if item == 'L' {
			num += 50
		} else if item == 'V' {
			num += 5
		} else if item == 'I' {
			if i+1 < bytesLen && bytes[i+1] == 'X' {
				i++
				num += 9
			} else if i+1 < bytesLen && bytes[i+1] == 'V' {
				i++
				num += 4
			} else {
				num += 1
			}
		}
	}

	return num
}

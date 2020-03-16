package tools

import "fmt"

func PrintSlice(slice interface{}) {
	switch s := slice.(type) {
	case []int:
		_int(s)
	case [][]int:
		__int(s)
	case []byte:
		_byte(s)
	case [][]byte:
		__byte(s)
	case []string:
		_string(s)
	case []bool:
		_bool(s)
	case [][]bool:
		__bool(s)
	default:
		panic("don't have this case.")
	}
}

func __int(v [][]int) {
	for i := range v {
		_int(v[i])
	}
	fmt.Println()
}

func _int(v []int) {
	for i := range v {
		fmt.Printf("%d\t", v[i])
	}
	fmt.Println()
}

func __byte(v [][]byte) {
	for i := range v {
		_byte(v[i])
	}
	fmt.Println()
}

func _byte(v []byte) {
	for i := range v {
		fmt.Printf("%c\t", v[i])
	}
	fmt.Println()
}

func _string(v []string) {
	for i := range v {
		fmt.Printf("%s\n", v[i])
	}
	fmt.Println()
}

func __bool(v [][]bool) {
	for i := range v {
		_bool(v[i])
	}
	fmt.Println()
}

func _bool(v []bool) {
	for i := range v {
		if v[i] {
			fmt.Printf("√\t")
		} else {
			fmt.Printf("×\t")
		}

	}
	fmt.Println()
}

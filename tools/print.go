package tools

import "fmt"

// PrintSlice 打印Slice
func PrintSlice(slice interface{}) {
	switch s := slice.(type) {
	case []int:
		_Int(s)
	case [][]int:
		__Int(s)
	case []byte:
		_Byte(s)
	case [][]byte:
		__Byte(s)
	case []string:
		_String(s)
	case []bool:
		_Bool(s)
	case [][]bool:
		__Bool(s)
	default:
		panic("don't have this case.")
	}
}

func __Int(v [][]int) {
	for i := range v {
		_Int(v[i])
	}
	fmt.Println()
}

func _Int(v []int) {
	for i := range v {
		fmt.Printf("%d\t", v[i])
	}
	fmt.Println()
}

func __Byte(v [][]byte) {
	for i := range v {
		_Byte(v[i])
	}
	fmt.Println()
}

func _Byte(v []byte) {
	for i := range v {
		fmt.Printf("%c\t", v[i])
	}
	fmt.Println()
}

func _String(v []string) {
	for i := range v {
		fmt.Printf("%s\n", v[i])
	}
	fmt.Println()
}

func __Bool(v [][]bool) {
	for i := range v {
		_Bool(v[i])
	}
	fmt.Println()
}

func _Bool(v []bool) {
	for i := range v {
		if v[i] {
			fmt.Printf("√\t")
		} else {
			fmt.Printf("×\t")
		}

	}
	fmt.Println()
}

package tools

import "fmt"

// PrintSlice 打印Slice
func PrintSlice(slice interface{}) {
	switch s := slice.(type) {
	case []int:
		_Ints(s)
	case [][]int:
		_Int(s)
	case []byte:
		_Bytes(s)
	case [][]byte:
		_Byte(s)
	case []string:
		_String(s)
	case []bool:
		_Bools(s)
	case [][]bool:
		_Bool(s)
	default:
		panic("don't have this case.")
	}
}

func _Int(v [][]int) {
	for i := range v {
		_Ints(v[i])
	}
	fmt.Println()
}

func _Ints(v []int) {
	for i := range v {
		fmt.Printf("%d\t", v[i])
	}
	fmt.Println()
}

func _Byte(v [][]byte) {
	for i := range v {
		_Bytes(v[i])
	}
	fmt.Println()
}

func _Bytes(v []byte) {
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

func _Bool(v [][]bool) {
	for i := range v {
		_Bools(v[i])
	}
	fmt.Println()
}

func _Bools(v []bool) {
	for i := range v {
		if v[i] {
			fmt.Printf("√\t")
		} else {
			fmt.Printf("×\t")
		}

	}
	fmt.Println()
}

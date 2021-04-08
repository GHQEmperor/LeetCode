package leetcode_test

import "testing"

type Test struct {
	Input  interface{}
	Output interface{}
}

func TestLeetCode(t *testing.T) {
	tests := []Test{
		{},
	}

	for _, test := range tests {
		out := LeetCode(test.Input)
		if out != test.Output {
			t.Errorf("test %v\t, out %v\n", test, out)
		}
	}
}

func LeetCode(input interface{}) interface{} {
	return nil
}

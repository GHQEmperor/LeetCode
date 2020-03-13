package main

import (
	"fmt"
)

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	result := addTwoNumbers(&l1, &l2)
	p := result
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}

// l1 2->4->3  342
// l2 5->6->4  465
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var nextPlus int
	var result *ListNode
	for {
		var (
			l1Value, l2Value int
			l1tail, l2tail   bool
		)
		l1Value, l1tail, l1 = getValue(l1)
		l2Value, l2tail, l2 = getValue(l2)

		value := l1Value + l2Value + nextPlus
		nextPlus = value / 10
		value = value % 10
		//fmt.Println(nextPlus)
		if !(l1tail && l2tail && value == 0) {
			result = addValue(result, value)
		}

		if l1tail && l2tail {
			return result
		}

	}
}

func getValue(l *ListNode) (val int, tail bool, next *ListNode) {
	if l == nil {
		return 0, true, nil
	}
	return l.Val, false, l.Next
}

func addValue(r *ListNode, val int) (resp *ListNode) {
	if r == nil {
		r = &ListNode{
			Val:  val,
			Next: nil,
		}
		return r
	}
	p := r
	for {
		if p.Next != nil {
			p = p.Next
			continue
		}
		p.Next = &ListNode{
			Val:  val,
			Next: nil,
		}
		return r
	}
}

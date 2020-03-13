package main

import "fmt"

func main() {
	l := mergeTwoLists(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		}}, &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	})

	for l != nil {
		fmt.Println(l)
		l = l.Next
	}
}

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

/*
21. 合并两个有序链表
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l1, l2 = l2, l1
	}

	p1, p2 := l1, l2
	for p1 != nil && p2 != nil {
		if p1.Next == nil {
			p1.Next = p2
			break
		}
		if p1.Next.Val > p2.Val {
			node := p2
			p2 = p2.Next
			node.Next = p1.Next
			p1.Next = node
			p1 = p1.Next
		} else {
			p1 = p1.Next
		}
	}
	return l1
}

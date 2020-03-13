package main

import "fmt"

func main() {
	l := swapPairs(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	})

	for l != nil {
		fmt.Println(l.Val)
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
24. 两两交换链表中的节点
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例:
给定 1->2->3->4, 你应该返回 2->1->4->3.
*/

func swapPairs(head *ListNode) *ListNode {
	p1, p2, p3 := head, head, head
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	if head.Next.Next == nil {
		head.Next.Next = head
		head = head.Next
		head.Next.Next = nil
	}
	for i := 0; i < 2; i++ {
		// 双节点的情况已经剔除
		if p3.Next == nil {
			return head
		}
		p3 = p3.Next
	}
	p2 = p2.Next

	head = head.Next
	p2.Next = p1
	p1.Next = p3

	p2 = p2.Next // 最前面的标记
	p1 = p1.Next // 中间的标记
	p3 = p3.Next // 最后的标记
	for p3 != nil {
		temp := p3.Next
		p3.Next = p1
		p1.Next = temp
		p2.Next = p3
		p1, p3 = p3, p1
		for i := 0; i < 2; i++ {
			if p3.Next == nil {
				//if i == 1 {
				//	p1.Next.Next = p1
				//	p2.Next = p1.Next
				//	p1.Next = nil
				//}

				return head
			}
			p1 = p1.Next
			p2 = p2.Next
			p3 = p3.Next
		}
	}
	return head
}

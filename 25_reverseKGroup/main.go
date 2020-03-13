package main

import "fmt"

func main() {
	l := reverseKGroup(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}, 2)

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

/*
25. K 个一组翻转链表
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
k 是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

示例 :
给定这个链表：1->2->3->4->5
当 k = 2 时，应当返回: 2->1->4->3->5
当 k = 3 时，应当返回: 3->2->1->4->5
说明 :
你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	var result ListNode
	var pre *ListNode
	pre = &result

	for head != nil {
		start := head
		for i := 0; i < k; i++ {
			if head == nil {
				pre.Next = start
				return result.Next
			}
			if i == k {
				//fmt.Println(i, k)
				pre.Next = reverse(start, k)
				return result.Next
			}
			head = head.Next
		}
		pre.Next = reverse(start, k)
		pre = start
	}

	return result.Next
}

func reverse(head *ListNode, k int) *ListNode {
	var list ListNode
	var p *ListNode
	for i := 0; i < k; i++ {
		list.Next = head
		head = head.Next
		list.Next.Next = p
		p = list.Next
	}
	return list.Next
}

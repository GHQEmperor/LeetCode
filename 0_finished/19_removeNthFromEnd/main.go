package main

import "fmt"

func main() {
	l := removeNthFromEnd(&ListNode{
		Val: 1,
		//Next: nil,
		Next: &ListNode{
			Val: 2,
			//	Next: &ListNode{
			//		Val: 3,
			//		Next: &ListNode{
			//			Val: 4,
			//			Next: &ListNode{
			//				Val:  5,
			//				Next: nil,
			//			},
			//		},
			//	},
		},
	}, 1)
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
19. 删除链表的倒数第N个节点
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：
给定一个链表: 1->2->3->4->5, 和 n = 2.
当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：
给定的 n 保证是有效的。
进阶：
你能尝试使用一趟扫描实现吗？
todo: 记录此时的节点的前第 N+1位
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var (
		p, temp *ListNode
		count   int
	)

	p = head
	temp = head
	for p != nil {
		if count > n {
			temp = temp.Next
		}
		p = p.Next
		count++
	}

	if count == 1 {
		if n == 1 {
			return nil
		} else {
			return head
		}
	}

	if count == n {
		return head.Next
	}
	temp.Next = temp.Next.Next

	return head
}

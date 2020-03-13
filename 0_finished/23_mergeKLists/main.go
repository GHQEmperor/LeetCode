package main

import "fmt"

func main() {
	l := mergeKLists([]*ListNode{
		{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
		{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
		//nil, nil,
		{
			Val: 2,
			Next: &ListNode{
				Val:  6,
				Next: nil,
			},
		},
	})

	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}
	all := lists
	for {
		n = len(all)
		if n == 1 {
			return all[0]
		}
		var temp []*ListNode
		if n%2 == 1 {
			temp = append(temp, all[n-1])
		}
		for i := 0; i < n-1; i += 2 {
			temp = append(temp, mergeTwoLists(all[i], all[i+1]))
		}
		all = temp
	}

}

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

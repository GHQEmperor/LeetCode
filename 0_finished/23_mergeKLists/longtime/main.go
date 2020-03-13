package longtime

import (
	"fmt"
	"sort"
)

func main() {
	//	1->4->5,
	//	1->3->4,
	//	2->6
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

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 1 {
		return lists[0]
	}
	var res, p *ListNode
	//var changed bool
	for {
		for i := 0; i < n; i++ {
			//fmt.Println(i, n)
			if lists[i] == nil {
				lists[i], lists[n-1] = lists[n-1], lists[i]
				lists = lists[:n-1]
				i--
				n--
				//changed = true
			}
		}
		if len(lists) == 0 {
			break
		}
		//_selectI := 1
		if n <= 1 {
			if p != nil {
				p.Next = lists[0]
			}
			//_selectI = 0
		}
		//if lists[0].Val > lists[1].Val {
		sort.Slice(lists, func(i, j int) bool {
			return lists[i].Val < lists[j].Val
		})
		//changed = false
		//_selectI = 0
		//}

		if res == nil {
			res = lists[0]
			p = res
			lists[0] = lists[0].Next
			res.Next = nil
		} else if p != nil {
			p.Next = lists[0]
			lists[0] = lists[0].Next
			p = p.Next
			p.Next = nil
			//p = p.Next
		}

		//for _, v := range lists {
		//	fmt.Println(v)
		//}
		//return nil
	}
	return res
}

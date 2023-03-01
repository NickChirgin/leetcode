package linkedlist

import "sort"

func mergeKLists(lists []*ListNode) *ListNode {
	res := []int{}
	if len(lists) == 0 || lists == nil{
			return nil
	}
	for _, head := range lists {
			c := head
			for c != nil {
					res = append(res, c.Val)
					c = c.Next
			}
	}
	if len(res) == 0 {
			return nil
	}
	sort.Ints(res)
	head := &ListNode{Val: res[0]}
	if len(res) == 1 {
			return head
	}
	curr := &ListNode{}
	head.Next = curr
	for i, _ := range res {
			if i +1 < len(res) {
					curr.Val = res[i+1]
					if i + 2 == len(res) {
							break
					}
					curr.Next= &ListNode{}
					curr = curr.Next
			}
	}
	if head.Next == nil {
			return nil
	}
	return head
}
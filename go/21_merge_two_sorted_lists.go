package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := ListNode{Val: 0}
	l3 := &dummyHead

	// Use two independent pointers to traverse list1 and list2.
	// Append the smaller of the two values to the result list.

	for l1 != nil || l2 != nil {
		if l1 != nil {
			if l2 != nil {
				if l1.Val > l2.Val {
					l3.Next = &ListNode{Val: l2.Val}
					l3 = l3.Next
					l2 = l2.Next
				} else {
					l3.Next = &ListNode{Val: l1.Val}
					l3 = l3.Next
					l1 = l1.Next
				}
			} else {
				l3.Next = &ListNode{Val: l1.Val}
				l3 = l3.Next
				l1 = l1.Next
			}
		} else {
			// means l2 is not nil
			l3.Next = &ListNode{Val: l2.Val}
			l3 = l3.Next
			l2 = l2.Next
		}
	}

	return dummyHead.Next
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	x := mergeTwoLists(l1, l2)
	for x != nil {
		fmt.Print(x.Val, " ")
		x = x.Next
	}
}

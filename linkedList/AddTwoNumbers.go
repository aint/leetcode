package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	r := 0
	var head, l3 *ListNode
	for true {
		if l1 == nil && l2 == nil {
			if r == 1 {
				l3.Next = &ListNode{1, nil}
			}
			break
		}

		v1 := 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		v2 := 0
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		v3 := v1 + v2 + r
		if v3 >= 10 {
			v3 = v3 % 10
			r = 1
		} else {
			r = 0
		}

		if l3 == nil {
			head = &ListNode{v3, nil}
			l3 = head
		} else {
			l3.Next = &ListNode{v3, nil}
			l3 = l3.Next
		}
	}

	return head
}

// 1 9 9
//   6 6
// 2 6 5  (r = 1)

//   9 9
//     3
// 1 0 2  (r = 1)

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	first, next := head, head.Next

	for next != nil {
		if first.Val == next.Val {
			first.Next = next.Next
		} else {
			first = next
		}
		next = next.Next
	}
	return head
}

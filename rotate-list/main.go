package main

import "log"

func main() {
	root := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	root = rotateRight(root, 1)

	for root != nil {
		log.Println(root.Val)
		root = root.Next
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	// 建立父节点hash表
	next := head
	count := 0
	parents := make(map[*ListNode]*ListNode)
	for next != nil {
		if next.Next != nil {
			parents[next.Next] = next
		} else {
			parents[head] = next
		}
		next = next.Next
		count++
	}
	k = k % count

	for k != 0 {
		k--

		last := parents[head]
		lastParent := parents[last]

		lastParent.Next = nil
		last.Next = head
		head = last
	}
	return head
}

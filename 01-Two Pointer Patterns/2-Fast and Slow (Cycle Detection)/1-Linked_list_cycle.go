package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	seen := make(map[*ListNode]bool)

	for head.Next != nil {
		if exist := seen[head]; exist {
			return true
		}

		seen[head] = true

		head = head.Next
	}

	return false
}

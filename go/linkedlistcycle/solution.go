package linkedlistcycle

type ListNode struct {
	Val  int
	Next *ListNode
}

// source: https://leetcode.com/problems/linked-list-cycle
func hasCycle(head *ListNode) bool {
	slow := head
	fast := slow

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

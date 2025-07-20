package linkedlistcycletwo

type ListNode struct {
	Val  int
	Next *ListNode
}

// source: https://leetcode.com/problems/linked-list-cycle-ii
func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			for fast != head {
				fast = fast.Next
				head = head.Next
			}

			return fast
		}
	}

	return nil
}

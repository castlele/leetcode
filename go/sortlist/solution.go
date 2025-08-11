package sortlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// source: https://leetcode.com/problems/sort-list
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	second := split(head)

	head = sortList(head)
	second = sortList(second)

	return mergeLists(head, second)
}

func split(head *ListNode) *ListNode {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next

		if fast != nil {
			slow = slow.Next
		}
	}

	second := slow.Next
	slow.Next = nil

	return second
}

func mergeLists(lhs, rhs *ListNode) *ListNode {
	if lhs == nil {
		return rhs
	} else if rhs == nil {
		return lhs
	}

	if lhs.Val < rhs.Val {
		lhs.Next = mergeLists(lhs.Next, rhs)
		return lhs
	} else {
		rhs.Next = mergeLists(lhs, rhs.Next)
		return rhs
	}
}

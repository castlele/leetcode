package rotatelist

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	lhs := head
	rhs := 0
	n := 0

	for lhs != nil {
		n++
		lhs = lhs.Next
	}

	if n == 1 {
		return head
	}

	k %= n
	rhs = (n - 1) - k
	lhs = head

	for rhs != n-1 {
		rhs++
		lhs = lhs.Next
	}

	newHead := lhs.Next
	lhs.Next = nil
	lhs = newHead

	for lhs.Next != nil {
		lhs = lhs.Next
	}

	lhs.Next = head

	return newHead
}

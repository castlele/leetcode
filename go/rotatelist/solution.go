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
	rhs := head
	n := 0

	for lhs != nil {
		n++
		lhs = lhs.Next
	}

	if n == 1 {
		return head
	}

	k %= n

	if k == 0 {
		return head
	}

	n = 0
	lhs = head

	for n < k {
		n++
		rhs = rhs.Next
	}

	for rhs.Next != nil {
		lhs = lhs.Next
		rhs = rhs.Next
	}

	newHead := lhs.Next
	lhs.Next = nil
	rhs.Next = head

	return newHead
}

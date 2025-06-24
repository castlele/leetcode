package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

// source: https://leetcode.com/problems/add-two-numbers/description/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var root *ListNode
	var cur *ListNode
	mind := 0
	lhs := l1
	rhs := l2

	for lhs != nil && rhs != nil {
		sum := lhs.Val + rhs.Val + mind
		newNode := &ListNode{Val: sum % 10}

		lhs = lhs.Next
		rhs = rhs.Next

		if root == nil {
			root = newNode
			cur = root
		} else {
			cur.Next = newNode
			cur = cur.Next
		}

		mind = sum / 10
	}

	for lhs != nil {
		sum := lhs.Val + mind
		lhs = lhs.Next
		newNode := &ListNode{Val: sum % 10}
		mind = sum / 10
		cur.Next = newNode
		cur = cur.Next
	}

	for rhs != nil {
		sum := rhs.Val + mind
		rhs = rhs.Next
		newNode := &ListNode{Val: sum % 10}
		mind = sum / 10
		cur.Next = newNode
		cur = cur.Next
	}

	if mind > 0 {
		cur.Next = &ListNode{Val: mind}
	}

	return root
}

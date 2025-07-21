package palindromelinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// source: https://leetcode.com/problems/palindrome-linked-list
func isPalindrome(head *ListNode) bool {
	var stack []int
	slow := head
	fast := slow

	for fast != nil && fast.Next != nil {
		stack = append(stack, slow.Val)
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 1, 2, 3, 3, 2, 1
	//          ^        ^
	//          |        |
	//          +-slow   +-fast
	if slow != nil && fast == nil {
		top := stack[len(stack)-1]

		if top != slow.Val {
			return false
		}

		stack = stack[:len(stack)-1]
	}

	for slow.Next != nil {
		slow = slow.Next
		top := stack[len(stack)-1]

		if top != slow.Val {
			return false
		}

		stack = stack[:len(stack)-1]
	}

	return true
}

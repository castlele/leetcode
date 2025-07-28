package removeduplicatesfromsortedlisttwo

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{
		Next: head,
	}

	left, right := dummy, head.Next
	counter := 1

	for right != nil {
		if left.Next.Val != right.Val {
			if counter != 1 {
				left.Next = right
				counter = 1
			} else {
				left = left.Next
			}
		} else {
			counter++
		}

		right = right.Next
	}

	if counter != 1 {
		left.Next = nil
	}

	return dummy.Next
}

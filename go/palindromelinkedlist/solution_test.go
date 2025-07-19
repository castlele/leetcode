package palindromelinkedlist

import (
	"fmt"
	"testing"
)

func TestPalindromeLinkedList(t *testing.T) {
	cases := []struct {
		input  []int
		output bool
	}{
		{[]int{1, 2, 2, 1}, true},
		{[]int{1, 2, 3, 4, 3, 2, 1}, true},
		{[]int{1, 2}, false},
	}

	for i, ts := range cases {
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			t.Parallel()
			res := isPalindrome(createLinkedList(ts.input))

			if ts.output != res {
				t.Errorf("Expected: %v; got: %v", ts.output, res)
			}
		})
	}
}

func createLinkedList(arr []int) *ListNode {
	root := &ListNode{Val: arr[0]}
	cur := root

	for i := 1; i < len(arr); i++ {
		val := arr[i]
		node := &ListNode{Val: val}

		cur.Next = node
		cur = node
	}

	return root
}

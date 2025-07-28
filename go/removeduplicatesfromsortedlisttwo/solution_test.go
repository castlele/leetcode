package removeduplicatesfromsortedlisttwo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveDuplicatesFromSortedListTwo(t *testing.T) {
	cases := []struct {
		input  []int
		output []int
	}{
		{[]int{1, 2, 3, 3, 4, 4, 5}, []int{1, 2, 5}},
		// {[]int{1, 1, 1, 2, 3}, []int{2, 3}},
	}

	for i, ts := range cases {
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			t.Parallel()
			res := deleteDuplicates(createLinkedList(ts.input))

			if !reflect.DeepEqual(ts.output, unwrapLinkedList(res)) {
				t.Errorf("Expected: %v; got: %v", ts.output, unwrapLinkedList(res))
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

func unwrapLinkedList(root *ListNode) []int {
	var res []int
	cur := root

	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}

	return res
}

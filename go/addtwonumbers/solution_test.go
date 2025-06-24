package addtwonumbers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	cases := []struct {
		l1     []int
		l2     []int
		output []int
	}{
		{l1: []int{2, 4, 3}, l2: []int{5, 6, 4}, output: []int{7, 0, 8}},
		{l1: []int{0}, l2: []int{0}, output: []int{0}},
		{l1: []int{9, 9, 9, 9, 9, 9, 9}, l2: []int{9, 9, 9, 9}, output: []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}

	for i, ts := range cases {
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			res := addTwoNumbers(createLinkedList(ts.l1), createLinkedList(ts.l2))
			unwrappedRes := unwrapLinkedList(res)

			if !reflect.DeepEqual(ts.output, unwrappedRes) {
				t.Errorf("Expected: %v; got: %v", ts.output, unwrappedRes)
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

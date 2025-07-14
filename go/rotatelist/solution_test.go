package rotatelist

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRotateList(t *testing.T) {
	cases := []struct {
		root   *ListNode
		k      int
		output []int
	}{
		{createLinkedList([]int{1, 2, 3, 4, 5}), 2, []int{4, 5, 1, 2, 3}},
		{createLinkedList([]int{0, 1, 2}), 4, []int{2, 0, 1}},
	}

	for i, ts := range cases {
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			t.Parallel()
			res := rotateRight(ts.root, ts.k)

			if !reflect.DeepEqual(ts.output, unwrapLinkedList(res)) {
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

func unwrapLinkedList(root *ListNode) []int {
	var res []int
	cur := root

	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}

	return res
}

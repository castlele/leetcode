package sortlist

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestSortList(t *testing.T) {
	cases := []struct {
		input []int
	}{
		{[]int{4, 2, 1, 3}},
		{[]int{-1, 5, 3, 4, 0}},
	}

	for i, ts := range cases {
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			t.Parallel()
			output := append([]int{}, ts.input...)
			sort.Ints(output)

			res := unwrapLinkedList(sortList(createLinkedList(ts.input)))

			if !reflect.DeepEqual(output, res) {
				t.Errorf("Expected: %v; got: %v", output, res)
			}
		})
	}
}

func createLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

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

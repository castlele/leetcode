package linkedlistcycletwo

import (
	"fmt"
	"testing"
)

type TS struct {
	head   *ListNode
	output *ListNode
}

func TestLinkedListCycle(t *testing.T) {
	cases := []TS{
		createTS_1(),
		createTS_2(),
	}

	for i, ts := range cases {
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			t.Parallel()
			res := detectCycle(ts.head)

			if res != ts.output {
				t.Errorf("Expected: %v; got: %v", ts.output, res)
			}
		})
	}
}

func createTS_1() TS {
	head, tail := createLinkedList([]int{1, 2, 3, 4, 5})

	tail.Next = head

	return TS{head, head}
}

func createTS_2() TS {
	head, _ := createLinkedList([]int{1, 2, 3, 4, 5})

	return TS{head, nil}
}

func createLinkedList(arr []int) (*ListNode, *ListNode) {
	root := &ListNode{Val: arr[0]}
	cur := root

	for i := 1; i < len(arr); i++ {
		val := arr[i]
		node := &ListNode{Val: val}

		cur.Next = node
		cur = node
	}

	return root, cur
}

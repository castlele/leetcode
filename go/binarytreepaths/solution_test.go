package binarytreepaths

import (
	"fmt"
	"reflect"
	"testing"
)

const NoValue = -404

func TestBinaryTreePaths(t *testing.T) {
	cases := []struct {
		input  []int
		output []string
	}{
		{[]int{1, 2, 3, NoValue, 5}, []string{"1->2->5", "1->3"}},
		{[]int{1}, []string{"1"}},
	}

	for i, ts := range cases {
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			t.Parallel()
			root := createBinaryTree(ts.input)

			res := binaryTreePaths(root)

			if !reflect.DeepEqual(ts.output, res) {
				t.Errorf("Expected: %v; got: %v", ts.output, res)
			}
		})
	}
}

func createBinaryTree(arr []int) *TreeNode {
	if len(arr) == 0 || arr[0] == NoValue {
		return nil
	}

	root := &TreeNode{Val: arr[0]}
	queue := []*TreeNode{root}
	i := 1

	for i < len(arr) {
		node := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(arr) && arr[i] != NoValue {
			node.Left = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Left)
		}
		i++

		// Right child
		if i < len(arr) && arr[i] != NoValue {
			node.Right = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

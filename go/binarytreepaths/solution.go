package binarytreepaths

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	var res []string
	var backtracking func(path string, cur *TreeNode)

	backtracking = func(path string, cur *TreeNode) {
		if cur.Left == nil && cur.Right == nil {
			res = append(res, strings.Clone(path))
			return
		}

		if cur.Left != nil {
			backtracking(path+fmt.Sprintf("->%d", cur.Left.Val), cur.Left)
		}

		if cur.Right != nil {
			backtracking(path+fmt.Sprintf("->%d", cur.Right.Val), cur.Right)
		}
	}

	backtracking(strconv.Itoa(root.Val), root)

	return res
}

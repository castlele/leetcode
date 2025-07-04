package searchtwodmatrix

import "fmt"

// source: https://leetcode.com/problems/search-a-2d-matrix
func searchMatrix(matrix [][]int, target int) bool {
	var row []int
	lhs := 0
	rhs := len(matrix) - 1

	for lhs <= rhs {
		mid := lhs + (rhs-lhs)/2

		fmt.Println(matrix[mid][0], matrix[mid][len(matrix[mid])-1], target)

		if matrix[mid][0] <= target && matrix[mid][len(matrix[mid])-1] >= target {
			row = matrix[mid]
			break
		} else if matrix[mid][0] > target {
			rhs = mid - 1
		} else {
			lhs = mid + 1
		}
	}

	if len(row) == 0 {
		return false
	}

	lhs = 0
	rhs = len(row) - 1

	for lhs <= rhs {
		mid := lhs + (rhs-lhs)/2

		if row[mid] == target {
			return true
		} else if row[mid] > target {
			rhs = mid - 1
		} else {
			lhs = mid + 1
		}
	}

	return false
}

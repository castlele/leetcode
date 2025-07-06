package searchtwodmatrixtwo

// source: https://leetcode.com/problems/search-a-2d-matrix-ii
func searchMatrix(matrix [][]int, target int) bool {
	for rowIndex := range matrix {
		row := matrix[rowIndex]

		if row[0] <= target && row[len(row)-1] >= target {
			lhs := 0
			rhs := len(row) - 1

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
		}
	}

	return false
}

package subsetstwo

import "sort"

// source: https://leetcode.com/problems/subsets-ii
func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	var path []int
	var backtracking func(start int)

	sort.Ints(nums)

	backtracking = func(start int) {
		res = append(res, append([]int{}, path...))

		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			path = append(path, nums[i])
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtracking(0)

	return res
}

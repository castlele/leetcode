package permutationstwo

import (
	"sort"
)

// source: https://leetcode.com/problems/permutations-ii/description/
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	var res [][]int
	var backtrace func([]int)
	used := make([]bool, len(nums))

	backtrace = func(path []int) {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}

			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			used[i] = true
			path = append(path, nums[i])
			backtrace(path)
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrace([]int{})

	return res
}

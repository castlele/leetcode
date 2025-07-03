package findfirstandlastpositionofelementinsortedarray

// source: https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array
func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	lhs := 0
	rhs := len(nums) - 1

	for lhs <= rhs {
		mid := lhs + (rhs-lhs)/2

		if nums[mid] == target {
			res[0] = mid

			for l := mid; l >= 0 && nums[l] == target; l-- {
				res[0] = l
			}

			res[1] = mid
			for r := mid; r < len(nums) && nums[r] == target; r++ {
				res[1] = r
			}

			return res

		} else if target < nums[mid] {
			rhs = mid - 1
		} else {
			lhs = mid + 1
		}
	}

	return res
}

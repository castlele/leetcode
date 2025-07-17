package sortcolors

// source: https://leetcode.com/problems/sort-colors
func sortColors(nums []int) {
	lhs := 0
	rhs := 0

	for _, num := range nums {
		switch num {
		case 0:
			lhs++
		case 1:
			rhs++
		default:
		}
	}

	for i := range nums {
		if lhs > 0 {
			nums[i] = 0
			lhs--
		} else if rhs > 0 {
			nums[i] = 1
			rhs--
		} else {
			nums[i] = 2
		}
	}
}

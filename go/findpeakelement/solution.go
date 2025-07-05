package findpeakelement

// source: https://leetcode.com/problems/find-peak-element
func findPeakElement(nums []int) int {
	n := len(nums)

	if n == 1 {
		return 0
	} else if nums[0] > nums[1] {
		return 0
	} else if nums[n-1] > nums[n-2] {
		return n - 1
	}

	lhs := 1
	rhs := n - 2

	for lhs <= rhs {
		mid := lhs + (rhs-lhs)/2

		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		} else if nums[mid] < nums[mid-1] {
			rhs = mid - 1
		} else {
			lhs = mid + 1
		}
	}

	return -1
}

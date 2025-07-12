package peakindexinmountainarray

// source: https://leetcode.com/problems/peak-index-in-a-mountain-array
func peakIndexInMountainArray(arr []int) int {
	if len(arr) == 3 {
		return 1
	}

	lhs := 0
	rhs := len(arr) - 1

	for lhs <= rhs {
		mid := lhs + (rhs-lhs)/2

		if arr[mid - 1] < arr[mid] && arr[mid + 1] < arr[mid] {
			return mid
		} else if arr[mid - 1] >= arr[mid] {
			rhs = mid
		} else {
			lhs = mid
		}
	}

    return 0
}

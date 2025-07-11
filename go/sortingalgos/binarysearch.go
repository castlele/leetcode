package sortingalgos

import (
	"cmp"
	"fmt"
)

// source: https://leetcode.com/problems/binary-search
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	lhs := 0
	rhs := len(nums) - 1

	for lhs <= rhs {
		mid := lhs + (rhs-lhs)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] >= target {
			rhs = mid - 1
		} else {
			lhs = mid + 1
		}
	}

	return -1
}

func binarySearch[V cmp.Ordered](nums []V, target V) *Optional[int] {
	if len(nums) == 0 {
		return None[int]()
	}

	lhs := 0
	rhs := len(nums) - 1

	for lhs <= rhs {
		mid := lhs + (rhs-lhs)/2

		fmt.Println(mid)

		if nums[mid] == target {
			return Opt(mid)
		} else if target <= nums[mid] {
			rhs = mid - 1
		} else {
			lhs = mid + 1
		}
	}

	return None[int]()
}

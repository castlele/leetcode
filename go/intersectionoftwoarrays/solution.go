package intersectionoftwoarrays

import (
	"maps"
	"slices"
	"sort"
)

// source: https://leetcode.com/problems/intersection-of-two-arrays
func intersection(nums1 []int, nums2 []int) []int {
	res := make(map[int]bool)
	var iter []int
	var target []int

	if len(nums1) > len(nums2) {
		iter = nums2
		target = nums1
	} else {
		iter = nums1
		target = nums2
	}

	sort.Ints(target)

	for _, num := range iter {
		if val, ok := res[num]; ok && val {
			continue
		}

		lhs := 0
		rhs := len(target) - 1

		for lhs <= rhs {
			mid := lhs + (rhs-lhs)/2

			if target[mid] == num {
				res[num] = true
				break
			} else if num > target[mid] {
				lhs = mid + 1
			} else {
				rhs = mid - 1
			}
		}
	}

	return slices.Collect(maps.Keys(res))
}

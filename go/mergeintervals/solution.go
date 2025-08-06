package mergeintervals

import (
	"fmt"
	"reflect"
	"slices"
)

// sourse: https://leetcode.com/problems/merge-intervals
func Merge(intervals [][]int) [][]int {
	var res [][]int

	slices.SortFunc(intervals, func(a []int, b []int) int {
		//  0: a = b
		// -1: a < b
		//  1: a > b
		if a[1] == b[1] && a[0] == b[0] {
			return 0
		} else if a[0] < b[0] && a[1] < b[1] {
			return -1
		} else {
			return 1
		}
	})

	fmt.Println(intervals)

	i, j := 0, 1

	for j != len(intervals) {
		lhs := intervals[i]
		rhs := intervals[j]

		if rhs[0] >= lhs[0] && rhs[0] <= lhs[1] {
			res = append(res, []int{lhs[0], rhs[1]})
			i = j
			intervals[i] = []int{lhs[0], rhs[1]}
			j++
		} else {
			if reflect.DeepEqual(res[len(res)-1], lhs) {
				res = append(res, rhs)
			} else {
				res = append(res, lhs)
			}
			i = j
			j++
		}
	}

	if !reflect.DeepEqual(res[len(res)-1], intervals[i]) {
		res = append(res, intervals[i])
	}

	fmt.Println(res)

	return res
}

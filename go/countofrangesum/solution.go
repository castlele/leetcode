package countofrangesum

// source: https://leetcode.com/problems/count-of-range-sum
func countRangeSum(nums []int, lower int, upper int) int {
	prefix := make([]int, len(nums)+1)

	for i, num := range nums {
		prefix[i+1] = prefix[i] + num
	}

	var mergeSort func(lhs, rhs int) int

	mergeSort = func(lhs, rhs int) int {
		if rhs-lhs <= 1 {
			return 0
		}
		mid := (lhs + rhs) / 2
		count := mergeSort(lhs, mid) + mergeSort(mid, rhs)

		j, k := mid, mid
		for i := lhs; i < mid; i++ {
			for k < rhs && prefix[k]-prefix[i] < lower {
				k++
			}
			for j < rhs && prefix[j]-prefix[i] <= upper {
				j++
			}
			count += j - k
		}

		tmp := make([]int, rhs-lhs)
		p1, p2, idx := lhs, mid, 0
		for p1 < mid && p2 < rhs {
			if prefix[p1] <= prefix[p2] {
				tmp[idx] = prefix[p1]
				p1++
			} else {
				tmp[idx] = prefix[p2]
				p2++
			}
			idx++
		}
		for p1 < mid {
			tmp[idx] = prefix[p1]
			p1++
			idx++
		}
		for p2 < rhs {
			tmp[idx] = prefix[p2]
			p2++
			idx++
		}
		copy(prefix[lhs:rhs], tmp)

		return count
	}

	return mergeSort(0, len(prefix))
}

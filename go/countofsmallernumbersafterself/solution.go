package countofsmallernumbersafterself

type pair struct {
	value int
	index int
}

var (
	res []int
)

// source: https://leetcode.com/problems/count-of-smaller-numbers-after-self
func countSmaller(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	res = make([]int, len(nums))
	pairs := make([]pair, len(nums))

	for i, num := range nums {
		pairs[i] = pair{value: num, index: i}
	}

	sort(pairs)

	return res
}

func sort(nums []pair) {
	var mergeSort func(lhs, rhs int)
	var merge func(lhs, mid, rhs int)

	merge = func(lhs, mid, rhs int) {
		n1 := mid - lhs + 1
		n2 := rhs - mid
		i, j, k := 0, 0, lhs
		L := make([]pair, n1)
		R := make([]pair, n2)

		for ; i < n1; i++ {
			L[i] = nums[lhs+i]
		}

		for ; j < n2; j++ {
			R[j] = nums[mid+j+1]
		}

		i, j = 0, 0

		for i < n1 && j < n2 {
			if L[i].value > R[j].value {
				nums[k] = L[i]
				res[L[i].index] += n2 - j
				i++
			} else {
				nums[k] = R[j]
				j++
			}

			k++
		}

		for i < n1 {
			nums[k] = L[i]
			i++
			k++
		}

		for j < n2 {
			nums[k] = R[j]
			j++
			k++
		}
	}

	mergeSort = func(lhs, rhs int) {
		if lhs >= rhs {
			return
		}

		mid := lhs + (rhs-lhs)/2

		mergeSort(lhs, mid)
		mergeSort(mid+1, rhs)
		merge(lhs, mid, rhs)
	}

	mergeSort(0, len(nums)-1)
}

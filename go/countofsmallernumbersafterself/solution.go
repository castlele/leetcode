package countofsmallernumbersafterself

// source: https://leetcode.com/problems/count-of-smaller-numbers-after-self
func countSmaller(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	res := make([]int, len(nums))

	for i, num := range nums {
		for j := i+1; j < len(nums); j++ {
			if nums[j] < num {
				res[i]++
			}
		}
	}

	return res
}

func sort(nums []int) []int {
	copy := append([]int{}, nums...)

	var mergeSort func(lhs, rhs int)
	var merge func(lhs, mid, rhs int)

	merge = func(lhs, mid, rhs int) {
		n1 := mid - lhs + 1
		n2 := rhs - mid
		i, j, k := 0, 0, lhs
		L := make([]int, n1)
		R := make([]int, n2)

		for ; i < n1; i++ {
			L[i] = copy[lhs+i]
		}

		for ; j < n2; j++ {
			R[j] = copy[mid+j+1]
		}

		i, j = 0, 0

		for i < n1 && j < n2 {
			if L[i] < R[j] {
				copy[k] = L[i]
				i++
			} else {
				copy[k] = R[j]
				j++
			}

			k++
		}

		for i < n1 {
			copy[k] = L[i]
			i++
			k++
		}

		for j < n2 {
			copy[k] = R[j]
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

	mergeSort(0, len(copy)-1)

	return copy
}

func find(nums []int, num int) int {
	lhs := 0
	rhs := len(nums) - 1

	for lhs <= rhs {
		mid := lhs + (rhs-lhs)/2

		if nums[mid] == num {
			return mid
		} else if nums[mid] > num {
			rhs = mid - 1
		} else {
			lhs = mid + 1
		}
	}

	return -1
}

package sortarray

// source: https://leetcode.com/problems/sort-an-array
func sortArray(nums []int) []int {
	res := append([]int{}, nums...)

	mergeSort(res, 0, len(res)-1)

	return res
}

func mergeSort(nums []int, lhs, rhs int) {
	if lhs >= rhs {
		return
	}

	mid := lhs + (rhs-lhs)/2

	mergeSort(nums, lhs, mid)
	mergeSort(nums, mid+1, rhs)
	merge(nums, lhs, mid, rhs)
}

func merge(nums []int, lhs, mid, rhs int) {
	n1 := mid - lhs + 1
	n2 := rhs - mid
	i, j := 0, 0
	L := make([]int, n1)
	R := make([]int, n2)

	for ; i < n1; i++ {
		L[i] = nums[lhs+i]
	}

	for ; j < n2; j++ {
		R[j] = nums[mid+j+1]
	}

	i, j = 0, 0
	k := lhs

	for i < n1 && j < n2 {
		if L[i] < R[j] {
			nums[k] = L[i]
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

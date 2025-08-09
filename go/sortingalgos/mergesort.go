package sortingalgos

import (
	"cmp"
)

func MergeSort[V cmp.Ordered](arr []V) {
	var sort func(arr []V, lhs, rhs int)
	var merge func(arr []V, lhs, mid, rhs int)

	merge = func(arr []V, lhs, mid, rhs int) {
		n1 := mid - lhs + 1
		n2 := rhs - mid

		i := 0
		j := 0
		L := make([]V, n1)
		R := make([]V, n2)

		for i = range n1 {
			L[i] = arr[lhs+i]
		}

		for j = range n2 {
			R[j] = arr[mid+1+j]
		}

		i, j = 0, 0
		k := lhs

		for i < n1 && j < n2 {
			if L[i] <= R[j] {
				arr[k] = L[i]
				i++
			} else {
				arr[k] = R[j]
				j++
			}

			k++
		}

		for i < n1 {
			arr[k] = L[i]
			k++
			i++
		}

		for j < n2 {
			arr[k] = R[j]
			k++
			j++
		}
	}

	sort = func(arr []V, lhs, rhs int) {
		if lhs >= rhs {
			return
		}

		mid := (lhs + rhs) / 2

		sort(arr, lhs, mid)
		sort(arr, mid+1, rhs)
		merge(arr, lhs, mid, rhs)
	}

	sort(arr, 0, len(arr)-1)
}

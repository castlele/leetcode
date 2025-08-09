package main

import (
	"fmt"

	"github.com/castlele/leetcode/sorting"
)

func main() {
	arr := []int{38, 27, 50, 10, 43}
	sorting.MergeSort(arr)

	fmt.Println(arr)
}

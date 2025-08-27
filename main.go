package main

import (
	"fmt"
)

func main() {
	lhs := 0
	rhs := 3

	mid := lhs + (rhs - lhs)/2
	fmt.Printf("%d %d %d\n", lhs, mid, rhs)

	lhs = 0
	rhs = mid
	mid = lhs + (rhs - lhs)/2
	fmt.Printf("%d %d %d\n", lhs, mid, rhs)
}

package main

import (
	"fmt"

	"github.com/castlele/leetcode/go/searchinsertpos"
)

func main() {
	s := &searchinsertpos.Searcher{}

	fmt.Println(s.SearchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(s.SearchInsert([]int{1, 3, 5, 6}, 2))
}

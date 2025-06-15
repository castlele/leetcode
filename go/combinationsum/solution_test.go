package combinationsum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	cases := []struct{
		candidates []int
		target int
		output [][]int
	} {
		{[]int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}},
		{[]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{[]int{2}, 1, [][]int{}},
	}

	for index, ts := range cases {
		t.Run(fmt.Sprintf("test %v", index), func(t *testing.T) {
			result := combinationSum(ts.candidates, ts.target)

			if !reflect.DeepEqual(ts.output, result) && len(ts.output) != len(result) {
				t.Fatalf("Expected: %v; got: %v", ts.output, result)
			}
		})
	}
}

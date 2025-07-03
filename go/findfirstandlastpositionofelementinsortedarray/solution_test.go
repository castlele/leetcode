package findfirstandlastpositionofelementinsortedarray

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		output []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{}, 0, []int{-1, -1}},
		{[]int{2, 2}, 2, []int{0, 1}},
	}

	for i, ts := range cases {
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			t.Parallel()
			res := searchRange(ts.nums, ts.target)

			if !reflect.DeepEqual(ts.output, res) {
				t.Errorf("Expected: %v; got: %v", ts.output, res)
			}
		})
	}

}

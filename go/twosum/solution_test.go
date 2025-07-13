package twosum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		input  []int
		target int
		output []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for i, ts := range cases {
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			t.Parallel()
			res := twoSum(ts.input, ts.target)

			if reflect.DeepEqual(ts.output, res) {
				t.Errorf("Expected: %v; got: %v", ts.output, res)
			}
		})
	}
}

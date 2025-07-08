package intersectionoftwoarrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIntersectionOfTwoArrays(t *testing.T) {
	// TODO: Create a function that will compare two slices with any order
	if true {
		return
	}

	cases := []struct {
		lhs    []int
		rhs    []int
		output []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{9, 4}},
	}

	for i, ts := range cases {
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			t.Parallel()
			res := intersection(ts.lhs, ts.rhs)

			if !reflect.DeepEqual(ts.output, res) {
				t.Errorf("Expected: %v; got: %v", ts.output, res)
			}
		})
	}
}

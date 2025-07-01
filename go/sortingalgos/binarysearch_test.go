package sortingalgos

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	cases := []struct {
		input  []int
		target int
		output *Optional[int]
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 3, Opt(2)},
		{[]int{1, 2, 3, 4, 5, 6}, 1, Opt(0)},
		{[]int{1, 2, 3, 4, 5, 6}, 6, Opt(5)},
		{[]int{1, 2, 3, 4, 5, 6}, 5, Opt(4)},
		{[]int{1, 2, 3, 4, 5, 6}, -1, None[int]()},
		{[]int{1, 2, 3, 4, 5, 6}, 7, None[int]()},
	}

	for i, ts := range cases {
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			t.Parallel()

			res := binarySearch(ts.input, ts.target)

			if !ts.output.Equals(res) {
				t.Errorf("Expected: %v; got: %v", ts.output, res)
			}
		})
	}
}

package peakindexinmountainarray

import (
	"fmt"
	"testing"
)

func TestPeakIndexInMountainArray(t *testing.T) {
	cases := []struct {
		input  []int
		output int
	}{
		{[]int{0, 1, 0}, 1},
		{[]int{0, 2, 1, 0}, 1},
		{[]int{0, 10, 5, 2}, 1},
		{[]int{3, 4, 5, 1}, 2},
		{[]int{3, 5, 4, 3, 1}, 1},
	}

	for i, ts := range cases {
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			t.Parallel()
			res := peakIndexInMountainArray(ts.input)

			if ts.output != res {
				t.Errorf("Expected: %v; got: %v", ts.output, res)
			}
		})
	}
}

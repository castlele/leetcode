package findpeakelement

import (
	"fmt"
	"testing"
)

func TestFindPeakElement(t *testing.T) {
	cases := []struct {
		input  []int
		output int
	}{
		{[]int{1, 2, 3, 1}, 2},
		{[]int{1, 2, 1, 3, 5, 6, 4}, 5},
		{[]int{1, 2}, 1},
		{[]int{3, 4, 3, 2, 1}, 1},
	}

	for i, ts := range cases {
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			t.Parallel()
			res := findPeakElement(ts.input)

			if ts.output != res {
				t.Errorf("Expected: %v; got: %v", ts.output, res)
			}
		})
	}

}

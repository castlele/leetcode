package countofrangesum

import (
	"fmt"
	"testing"
)

func TestCountOfRangeSum(t *testing.T) {
	cases := []struct {
		input  []int
		l      int
		r      int
		output int
	}{
		{[]int{-2, 5, -1}, -2, 2, 3},
		{[]int{0}, 0, 0, 1},
		{[]int{0, 0}, 0, 0, 3},
		{[]int{2147483647, -2147483648, -1, 0}, -1, 0, 4},
	}

	for i, ts := range cases {
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			t.Parallel()
			res := countRangeSum(ts.input, ts.l, ts.r)

			if res != ts.output {
				t.Errorf("Expected: %v; got: %v", ts.output, res)
			}
		})
	}
}

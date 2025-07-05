package searchtwodmatrix

import (
	"fmt"
	"testing"
)

func TestSearch2DMatrix(t *testing.T) {
	cases := []struct {
		input  [][]int
		target int
		output bool
	}{
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3, true},
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13, false},
	}

	for i, ts := range cases {
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			t.Parallel()
			res := searchMatrix(ts.input, ts.target)

			if ts.output != res {
				t.Errorf("Expected: %v; got: %v", ts.output, res)
			}
		})
	}

}

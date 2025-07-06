package searchtwodmatrixtwo

import (
	"fmt"
	"testing"
)

func TestSearch2DMatrixTwo(t *testing.T) {
	cases := []struct {
		input  [][]int
		target int
		output bool
	}{
		{[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5, true},
		// {[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20, false},
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

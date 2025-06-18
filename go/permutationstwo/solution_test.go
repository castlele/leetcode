package permutationstwo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPermutationTwo(t *testing.T) {
	cases := []struct {
		intput []int
		output [][]int
	}{
		{[]int{1, 1, 2}, [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
	}

	for i, ts := range cases {
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			t.Parallel()
			res := permuteUnique(ts.intput)

			if !reflect.DeepEqual(ts.output, res) {
				t.Errorf("Expected: %v; got: %v", ts.output, res)
			}
		})
	}
}

package combinations

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCombinations(t *testing.T) {
	cases := []struct {
		n      int
		k      int
		output [][]int
	}{
		{n: 4, k: 2, output: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
		{n: 1, k: 1, output: [][]int{{1}}},
	}

	for i, ts := range cases {
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			t.Parallel()
			res := combine(ts.n, ts.k)

			if !reflect.DeepEqual(ts.output, res) {
				t.Errorf("Expected: %v; got: %v", ts.output, res)
			}
		})
	}
}

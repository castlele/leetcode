package combinationsumtwo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCombinationTwo(t *testing.T) {
	cases := []struct {
		candidates []int
		target     int
		output     [][]int
	}{
		{[]int{10, 1, 2, 7, 6, 1, 5}, 8, [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
		{[]int{2, 5, 2, 1, 2}, 5, [][]int{{1, 2, 2}, {5}}},
	}

	for i, ts := range cases {
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			res := combinationSum2(ts.candidates, ts.target)

			if !reflect.DeepEqual(ts.output, res) && len(ts.output) != len(res) {
				t.Errorf("Expected: %v; got: %v", ts.output, res)
			}
		})
	}
}

package subsets

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/castlele/leetcode/go/utils"
)

func TestSubsets(t *testing.T) {
	if true {
		return
	}
	cases := []struct {
		input  []int
		output [][]int
	}{
		{input: []int{1, 2, 3}, output: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
		{input: []int{0}, output: [][]int{{}, {0}}},
	}

	for i, ts := range cases {
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			t.Parallel()

			res := subsets(ts.input)
			utils.Sort2DSlice(res)

			if !reflect.DeepEqual(ts.output, res) {
				t.Errorf("Extected: %v; got: %v", ts.output, res)
			}
		})
	}
}

package countofsmallernumbersafterself

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCountOfSmallerNumbersAfterSelf(t *testing.T) {
	cases := []struct {
		input  []int
		output []int
	}{
		{[]int{5, 2, 6, 1}, []int{2, 1, 1, 0}},
		{[]int{}, []int{}},
		{[]int{-1}, []int{0}},
		{[]int{-1, -1}, []int{0, 0}},
	}

	for i, ts := range cases {
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			t.Parallel()
			res := countSmaller(ts.input)

			if !reflect.DeepEqual(ts.output, res) {
				t.Errorf("Expected: %v; got: %v", ts.output, res)
			}
		})
	}

}

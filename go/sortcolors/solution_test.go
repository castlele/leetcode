package sortcolors

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	cases := []struct {
		input  []int
		output []int
	}{
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{[]int{2, 0, 1}, []int{0, 1, 2}},
	}

	for i, ts := range cases {
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			t.Parallel()
			sortColors(ts.input)

			if !reflect.DeepEqual(ts.output, ts.input) {
				t.Errorf("Expected: %v; got: %v", ts.output, ts.input)
			}
		})
	}
}

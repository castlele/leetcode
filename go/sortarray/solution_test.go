package sortarray

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestSortArray(t *testing.T) {
	cases := []struct {
		input []int
	}{
		{[]int{5, 2, 3, 1}},
		{[]int{5, 1, 1, 2, 0, 0}},
	}

	for i, ts := range cases {
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			t.Parallel()
			output := append([]int{}, ts.input...)
			sort.Ints(output)

			res := sortArray(ts.input)

			if !reflect.DeepEqual(output, res) {
				t.Errorf("Expected: %v; got: %v", output, res)
			}
		})
	}
}

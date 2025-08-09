package sortingalgos

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestMergeSort(t *testing.T) {
	cases := []struct {
		input []int
	}{
		{[]int{38, 27, 43, 10}},
		{[]int{38, 27, 50, 10, 43}},
		{[]int{50, 43}},
		{[]int{}},
		{[]int{100}},
	}

	for i, ts := range cases {
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			t.Parallel()
			res := append([]int{}, ts.input...)
			sort.Ints(ts.input)

			MergeSort(res)

			if !reflect.DeepEqual(ts.input, res) {
				t.Errorf("Expected: %v; got: %v", ts.input, res)
			}
		})
	}
}

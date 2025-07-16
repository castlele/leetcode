package mergesortedarray

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeSortedArray(t *testing.T) {
	cases := []struct {
		n1     []int
		m      int
		n2     []int
		n      int
		output []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{1}, 1, []int{}, 0, []int{1}},
		{[]int{0}, 0, []int{1}, 1, []int{1}},
	}

	for i, ts := range cases {
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			t.Parallel()
			merge(ts.n1, ts.m, ts.n2, ts.n)

			if !reflect.DeepEqual(ts.output, ts.n1) {
				t.Errorf("Expected: %v; got: %v", ts.output, ts.n1)
			}
		})
	}
}

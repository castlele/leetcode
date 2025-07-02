package sqrtx

import (
	"fmt"
	"testing"
)

func TestSqrtx(t *testing.T) {
	cases := []struct {
		input  int
		output int
	}{
		{input: 4, output: 2},
		{input: 8, output: 2},
		{input: 1, output: 1},
		{input: 9, output: 3},
	}

	for i, ts := range cases {
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			t.Parallel()

			res := mySqrt(ts.input)

			if ts.output != res {
				t.Errorf("Expected: %v; got: %v", ts.output, res)
			}
		})
	}
}

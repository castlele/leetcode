package mycalendarone

import (
	"fmt"
	"testing"
)

func TestMyCalendarOne(t *testing.T) {
	sut := Constructor()

	cases := []struct {
		start  int
		end    int
		output bool
	}{
		{47, 50, true},
		{33, 41, true},
		{39, 45, false},
		{33, 42, false},
		{25, 32, true},
		{26, 35, false},
		{19, 25, true},
		{3, 8, true},
		{8, 13, true},
		{18, 27, false},
	}

	for i, ts := range cases {
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			res := sut.Book(ts.start, ts.end)

			if ts.output != res {
				t.Errorf("Expected: %v; got: %v", ts.output, res)
			}
		})
	}

}

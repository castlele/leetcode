package validpalindrome

import (
	"fmt"
	"testing"
)

func TestValidPalindrome(t *testing.T) {
	cases := []struct {
		input  string
		output bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
	}

	for i, ts := range cases {
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			t.Parallel()
			res := isPalindrome(ts.input)

			if ts.output != res {
				t.Errorf("Expected: %v; got: %v", ts.output, res)
			}
		})
	}

}

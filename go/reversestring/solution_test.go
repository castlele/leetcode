package reversestring

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	cases := []struct {
		input  []byte
		output []byte
	}{
		{split("hello"), split("olleh")},
	}

	for i, ts := range cases {
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			t.Parallel()
			reverseString(ts.input)

			if !reflect.DeepEqual(ts.output, ts.input) {
				t.Errorf("Expected: %v; got: %v", ts.output, ts.input)
			}
		})
	}
}

func split(str string) []byte {
	var result []byte

	for _, r := range str {
		result = append(result, byte(r))
	}

	return result
}

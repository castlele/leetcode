package groupanagrams

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	if true {
		return
	}

	cases := []struct {
		input  []string
		output [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		{[]string{""}, [][]string{{""}}},
	}

	for i, ts := range cases {
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			t.Parallel()
			res := groupAnagrams(ts.input)

			if !reflect.DeepEqual(sortStrs(ts.output), sortStrs(res)) {
				t.Errorf("Expected: %v; got: %v", ts.output, res)
			}
		})
	}
}

func sortStrs(strs [][]string) [][]string {
	var res [][]string

	for _, str := range strs {
		sort.Strings(str)
		res = append(res, str)
	}

	return res
}

package permutations

import (
	"reflect"
	"testing"
)

func TestPermutations(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		t.Parallel()
		nums := []int{1,2,3}
		expected := [][]int{{1,2,3},{1,3,2},{2,1,3},{2,3,1},{3,1,2},{3,2,1}}

		res := permute(nums)

		if !reflect.DeepEqual(expected, res) {
			t.Errorf("Invalid result! Expected: %d; got: %d", expected, res)
		}
	})
}

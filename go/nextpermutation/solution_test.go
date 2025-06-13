package nextpermutation

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		t.Parallel()
		nums := []int{1, 2, 3}
		expected := []int{1, 3, 2}

		nextPermutation(nums)

		if !reflect.DeepEqual(expected, nums) {
			t.Errorf("Invalid result! Expected: %v; got: %v\n", expected, nums)
		}
	})

	t.Run("test 2", func(t *testing.T) {
		t.Parallel()
		nums := []int{3, 2, 1}
		expected := []int{1, 2, 3}

		nextPermutation(nums)

		if !reflect.DeepEqual(expected, nums) {
			t.Errorf("Invalid result! Expected: %v; got: %v\n", expected, nums)
		}
	})

	t.Run("test 3", func(t *testing.T) {
		t.Parallel()
		nums := []int{1, 1, 5}
		expected := []int{1, 5, 1}

		nextPermutation(nums)

		if !reflect.DeepEqual(expected, nums) {
			t.Errorf("Invalid result! Expected: %v; got: %v\n", expected, nums)
		}
	})
}

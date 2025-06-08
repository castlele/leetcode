package plusone

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		t.Parallel()
		digits := []int{1, 2, 3}
		expected := []int{1, 2, 4}

		result := plusOne(digits)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Invalid result. Expected: %d; got: %d", expected, result)
		}
	})

	t.Run("test 2", func(t *testing.T) {
		t.Parallel()
		digits := []int{4, 3, 2, 1}
		expected := []int{4, 3, 2, 2}

		result := plusOne(digits)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Invalid result. Expected: %d; got: %d", expected, result)
		}
	})

	t.Run("test 3", func(t *testing.T) {
		t.Parallel()
		digits := []int{9}
		expected := []int{1, 0}

		result := plusOne(digits)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Invalid result. Expected: %d; got: %d", expected, result)
		}
	})

	t.Run("test 4", func(t *testing.T) {
		t.Parallel()
		digits := []int{9, 9, 9}
		expected := []int{1, 0, 0, 0}

		result := plusOne(digits)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Invalid result. Expected: %d; got: %d", expected, result)
		}
	})
}

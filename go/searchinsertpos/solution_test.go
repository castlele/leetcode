package searchinsertpos

import "testing"

func TestSearchInsert(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		t.Parallel()
		nums := []int{1, 3, 5, 6}
		target := 5
		expect := 2

		result := searchInsert(nums, target)

		if result != expect {
			t.Fatalf("Invalid result, expect: %d; got: %d", expect, result)
		}
	})
	t.Run("test 2", func(t *testing.T) {
		t.Parallel()
		nums := []int{1, 3, 5, 6}
		target := 2
		expect := 1

		result := searchInsert(nums, target)

		if result != expect {
			t.Fatalf("Invalid result, expect: %d; got: %d", expect, result)
		}
	})
	t.Run("test 3", func(t *testing.T) {
		t.Parallel()
		nums := []int{1, 3, 5, 6}
		target := 7
		expect := 4

		result := searchInsert(nums, target)

		if result != expect {
			t.Fatalf("Invalid result, expect: %d; got: %d", expect, result)
		}
	})
	t.Run("test 4", func(t *testing.T) {
		t.Parallel()
		nums := []int{1, 3}
		target := 2
		expect := 1

		result := searchInsert(nums, target)

		if result != expect {
			t.Fatalf("Invalid result, expect: %d; got: %d", expect, result)
		}
	})
}

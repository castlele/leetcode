package parselispexp

import "testing"

func TestParsingLispExp(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		t.Parallel()
		exp := "(let x 2 (mult x (let x 3 y 4 (add x y))))"
		expected := 14

		res := evaluate(exp)

		if res != expected {
			t.Errorf("Invalid result! Expected: %d; got: %d", expected, res)
		}
	})

	t.Run("test 2", func(t *testing.T) {
		t.Parallel()
		exp := "(let x 3 x 2 x)"
		expected := 2

		res := evaluate(exp)

		if res != expected {
			t.Errorf("Invalid result! Expected: %d; got: %d", expected, res)
		}
	})

	t.Run("test 3", func(t *testing.T) {
		t.Parallel()
		exp := "(let x 1 y 2 x (add x y) (add x y))"
		expected := 5

		res := evaluate(exp)

		if res != expected {
			t.Errorf("Invalid result! Expected: %d; got: %d", expected, res)
		}
	})

	t.Run("test 4", func(t *testing.T) {
		t.Parallel()
		exp := "(-1)"
		expected := -1

		res := evaluate(exp)

		if res != expected {
			t.Errorf("Invalid result! Expected: %d; got: %d", expected, res)
		}
	})

	t.Run("test 5", func(t *testing.T) {
		t.Parallel()
		exp := "(20)"
		expected := 20

		res := evaluate(exp)

		if res != expected {
			t.Errorf("Invalid result! Expected: %d; got: %d", expected, res)
		}
	})
}

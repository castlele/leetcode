package lenoflastword

import "testing"

func TestLenOfLastWord(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		t.Parallel()
		str := "Hello World"
		expected := 5

		res := lengthOfLastWord(str)

		if res != expected {
			t.Errorf("Invalid len! Expected: %d; got: %d", expected, res)
		}
	})

	t.Run("test 2", func(t *testing.T) {
		t.Parallel()
		str := "   fly me   to   the moon  "
		expected := 4

		res := lengthOfLastWord(str)

		if res != expected {
			t.Errorf("Invalid len! Expected: %d; got: %d", expected, res)
		}
	})

	t.Run("test 3", func(t *testing.T) {
		t.Parallel()
		str := "luffy is still joyboy"
		expected := 6

		res := lengthOfLastWord(str)

		if res != expected {
			t.Errorf("Invalid len! Expected: %d; got: %d", expected, res)
		}
	})
}

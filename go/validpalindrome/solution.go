package validpalindrome

import (
	"strings"
)

func isPalindrome(s string) bool {
	lhs := 0
	rhs := len(s) - 1

	isValidChar := func(s1 string) bool {
		if (s1 >= "A" && s1 <= "Z") || (s1 >= "a" && s1 <= "z") || (s1 >= "0" && s1 <= "9") {
			return true
		}
		return false
	}

	for lhs <= rhs {
		l := string(s[lhs])
		r := string(s[rhs])

		if !isValidChar(l) {
			lhs++
			continue
		}

		if !isValidChar(r) {
			rhs--
			continue
		}

		if strings.ToUpper(l) != strings.ToUpper(r) {
			return false
		}

		lhs++
		rhs--
	}

	return true
}

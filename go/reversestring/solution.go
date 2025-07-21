package reversestring

// source: https://leetcode.com/problems/reverse-string
func reverseString(s []byte) {
	lhs := 0
	rhs := len(s) - 1

	for lhs <= rhs {
		tmp := s[lhs]
		s[lhs] = s[rhs]
		s[rhs] = tmp

		lhs++
		rhs--
	}
}

package lenoflastword

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}

	i := len(s) - 1

	for i >= 0 && s[i] == ' ' {
		i--
	}

	if i < 0 {
		return 0
	}

	len := 0

	for i >= 0 && s[i] != ' ' {
		len++
		i--
	}

	return len
}

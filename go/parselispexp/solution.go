package parselispexp

import "strconv"

// source: https://leetcode.com/problems/parse-lisp-expression/description/
func evaluate(expression string) int {
	var s stack
	index := 0

	for index < len(expression) {
		token := expression[index]

		if token == '(' || token == ' ' || token == ')' {
			index++
			continue
		}

		_, err := strconv.Atoi(string(expression[index]))

		if expression[index] == '-' || err != nil {
			num, i := parseNum(index, expression)
			s = append(s, num)
			index = i
			continue
		}

		index++
		continue
		switch expression[index : index+3] {
		case "let":
		case "add":
		case "mul":
		}
	}

	return s.Evaluate()
}

func parseNum(index int, exp string) (string, int) {
	end := index + 1

	for end < len(exp) && exp[end] != ' ' && exp[end] != ')' && exp[end] != '(' {
		end++
	}

	return exp[index:end], end
}

type stack []string

func (s stack) Evaluate() int {
	el := s[len(s)-1]
	i, _ := strconv.Atoi(el)

	return i
}

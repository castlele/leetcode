package parselispexp

// source: https://leetcode.com/problems/parse-lisp-expression/description/
func evaluate(expression string) int {
	var s *stack
	index := 0

	for index < len(expression) {
		token := expression[index]

		if token == '(' {
			index++

			
		}
	}

	return s.Evaluate()
}

type stack []string

func (s *stack) Evaluate() int {
	return 0
}

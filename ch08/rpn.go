package ch08

import "strconv"

// ex 9.2
// Reverse Polish Notation
func RPN(expr []string) (int, error) {
	s := NewStack()

	isOperator := func(x string) bool {
		return x == "+" ||
			x == "-" ||
			x == "*" ||
			x == "/"
	}

	calculate := func(a, b int, ope string) int {
		switch ope {
		case "+":
			return a + b
		case "-":
			return a - b
		case "*":
			return a * b
		case "/":
			return a / b
		default:
			return 0
		}
	}

	for _, x := range expr {
		if isOperator(x) {
			b, err := s.Pop()
			if err != nil {
				return 0, err
			}
			a, err := s.Pop()
			if err != nil {
				return 0, err
			}
			if err := s.Push(calculate(a, b, x)); err != nil {
				return 0, err
			}
		} else {
			i, err := strconv.Atoi(x)
			if err != nil {
				return 0, err
			}
			if err := s.Push(i); err != nil {
				return 0, err
			}
		}
	}

	return s.Pop()
}

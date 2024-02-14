package calc

// Package for converting infix to postfix & evaling it

import (
	"fmt"
	"strconv"
)

func prec(s string) int {
	if (s == "/") || (s == "*") {
		return 2
	} else if (s == "+") || (s == "-") {
		return 1
	} else {
		return -1
	}
}
func isOp(r rune) bool {
	return (r == 42 || r == 43 || r == 45 || r == 47)
}

// Infix to postfix
func InfixTo(infix string) ([]string, error) {
	stack := NewStack[string]()
	var postfix []string

	for _, char := range infix {
		opchar := string(char)
		if char >= '0' && char <= '9' {
			postfix = append(postfix, opchar)

		} else if char == '(' {
			stack.Push(opchar)

		} else if char == ')' {
			for stack.Top() != "(" {
				postfix = append(postfix, stack.Top())
				stack.Pop()
			}
			stack.Pop()

		} else if isOp(char) {
			for !stack.IsEmpty() && prec(opchar) <= prec(stack.Top()) {
				postfix = append(postfix, stack.Top())
				stack.Pop()
			}
			stack.Push(opchar)
		} else {
			return postfix, fmt.Errorf("invalid operator: %c", char)
		}
	}
	for !stack.IsEmpty() {
		postfix = append(postfix, stack.Top())
		stack.Pop()
	}
	return postfix, nil
}

func Eval(postfix []string) (float64, error) {
	stack := make([]float64, 0)

	for _, token := range postfix {
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			stack = append(stack, num)
		} else {
			if len(stack) < 2 {
				return 0, fmt.Errorf("invalid postfix expression")
			}

			// Pop the last two operands
			op2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			op1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Perform the operation based on the operator
			switch token {
			case "+":
				stack = append(stack, op1+op2)
			case "-":
				stack = append(stack, op1-op2)
			case "*":
				stack = append(stack, op1*op2)
			case "/":
				if op2 == 0 {
					return 0, fmt.Errorf("division by zero")
				}

				stack = append(stack, op1/op2)
			default:
				return 0, fmt.Errorf("unknown operator: %s", token)
			}
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid postfix expression")
	}

	return stack[0], nil
}

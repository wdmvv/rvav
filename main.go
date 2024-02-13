package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isNumber(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}

func shuntingYard(infix string) (string, error) {
	out := make([]string, 0, len(infix))
	operators := map[rune]int{'*': 3, '/': 3, '+': 2, '-': 2}

	var stack []rune

	tokens := strings.FieldsFunc(infix, func(r rune) bool {
		return r == ' ' || r == '(' || r == ')'
	})

	for _, token := range tokens {
		if isNumber(token) {
			out = append(out, token)
		} else if len(token) == 1 && (token[0] == '(' || token[0] == ')') {
			r := rune(token[0])
			if r == '(' {
				stack = append(stack, r)
			} else {
				for len(stack) > 0 && stack[len(stack)-1] != '(' {
					out = append(out, string(stack[len(stack)-1]))
					stack = stack[:len(stack)-1]
				}
				if len(stack) == 0 {
					return "", fmt.Errorf("mismatched parentheses")
				}
				stack = stack[:len(stack)-1]
			}
		} else if len(token) == 1 && operators[rune(token[0])] > 0 {
			r := rune(token[0])
			for len(stack) > 0 && operators[stack[len(stack)-1]] >= operators[r] {
				out = append(out, string(stack[len(stack)-1]))
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, r)
		} else {
			return "", fmt.Errorf("invalid character or token: " + token)
		}
	}

	for len(stack) > 0 {
		if stack[len(stack)-1] == '(' || stack[len(stack)-1] == ')' {
			return "", fmt.Errorf("mismatched parentheses")
		}
		out = append(out, string(stack[len(stack)-1]))
		stack = stack[:len(stack)-1]
	}
	return strings.Join(out, " "), nil
}


func evaluateRPN(expression string) (float64, error) {
  tokens := strings.Fields(expression)
  stack := make([]float64, 0)

  for _, token := range tokens {
    if num, err := strconv.ParseFloat(token, 64); err == nil {
      stack = append(stack, num)
    } else {
      if len(stack) < 2 {
        return 0, fmt.Errorf("invalid RPN expression")
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
    return 0, fmt.Errorf("invalid RPN expression")
  }

  return stack[0], nil  
}

func main() {
	eq := "(2 + 2) * 2 / 3"
	rpn, err := shuntingYard(eq)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(evaluateRPN(rpn))
	}
}

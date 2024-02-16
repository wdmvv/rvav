package calc

// Package for converting infix to postfix & evaling it

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"orchestrator/config"
	"strconv"
	"time"
)

// Infix to postfix
func InfixToPostfix(infix string) ([]string, error) {
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

// calculating postfix with goroutines (future)
func Eval(postfix []string) (float64, error) {
	stack := NewStack[float64]()
	for _, token := range postfix {
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			stack.Push(num)
		} else {
			if stack.Len() < 2 {
				return 0, fmt.Errorf("invalid postfix expression")
			}
			op2 := stack.Pop()
			op1 := stack.Pop()

			topush, err := calcOp(op1, op2, token)
			if err != nil{
				return 0, err
			}
			stack.Push(topush)
		}
	}
	if stack.Len() != 1 {
		return 0, fmt.Errorf("invalid postfix expression")
	}
	return stack.Top(), nil
}

// calculate operation by sending request to agent
func calcOp(op1, op2 float64, sign string) (float64, error){
	timeout, err := config.Conf.SignTimeout(sign)
	if err != nil{
		return 0, err
	}
	send := AgentReqOut{op1, op2, sign, timeout}

	jsdata, _ := json.Marshal(send)
	url := fmt.Sprintf("http://127.0.0.1:%d/eval", config.Conf.AgentPort)
	client := &http.Client{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Conf.Timeout) * time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, bytes.NewBuffer(jsdata))
	if err != nil{
		return 0, fmt.Errorf("error creating request to agent")
	}
	
	resp, err := client.Do(req)
	if err != nil{
		if err == context.DeadlineExceeded{
			return 0, fmt.Errorf("request timeout")
		} else {
			fmt.Println(err)
			return 0, fmt.Errorf("failed to complete request")
		}
	}
	defer resp.Body.Close()

	var data AgentReqIn
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&data)
	if err != nil{
		return 0, fmt.Errorf("invalid response from agent")
	}
	
	if resp.StatusCode != http.StatusOK{
		return 0, fmt.Errorf(data.Errmsg)
	}

	return data.Result, nil	
}

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

type AgentReqOut struct {
	Op1     float64 `json:"op1"`
	Op2     float64 `json:"op2"`
	Sign    string  `json:"sign"`
	Timeout int     `json:"timeout"`
}

type AgentReqIn struct {
	Result float64 `json:"result"`
	Errmsg string  `json:"errmsg"`
}


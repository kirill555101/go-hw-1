package calc

import (
	"errors"
	"strconv"
	"strings"
	"./stack"
)

func GetPostfixNotation(str string) (string, error) {
	stack := stack.New()
	isNum := false
	result := ""

	for _, char := range str {
		if char > '0' && char < '9' {
			if isNum == false {
				result += " "
			}
			result += string(char)
			isNum = true

			continue
		}

		isNum = false

		switch char {
		case '(':
			stack.Push(char)
		case ')':
			for stack.Len() != 0 {
				popEl, _ := stack.Pop()
				if popEl == '(' {
					continue
				}
				result += " " + string(popEl.(rune))
			}
		case '+':
			fallthrough
		case '-':
			fallthrough
		case '*':
			fallthrough
		case '/':
			if stack.Len() == 0 {
				stack.Push(char)
				continue
			}
			peekEl, err := stack.Peek()
			if err != nil {
				return "", errors.New("invalid char")
			}

			if getPriorityOperation(peekEl.(rune)) < getPriorityOperation(char) {
				stack.Push(char)
				continue
			} else {
				popEl, _ := stack.Pop()
				result += " " + string(popEl.(rune))
				stack.Push(char)
			}

		default:
			return "", errors.New("invalid char")
		}
	}

	for stack.Len() != 0 {
		popEl, _ := stack.Pop()
		result += " " + string(popEl.(rune))
	}

	return result[1:], nil
}

func Calculate(inputString string) (float64, error) {
	str := strings.Split(inputString, " ")
	stack := stack.New()

	for _, val := range str {
		num, isNum := strconv.ParseFloat(val, 64)
		if isNum == nil {
			stack.Push(num)
			continue
		}

		lNum, lErr := stack.Pop()
		rNum, rErr := stack.Pop()

		if lErr != nil || rErr != nil {
			return 0, errors.New("Error in string")
		}

		res, err := getExpressionResult(lNum.(float64), rNum.(float64), val)
		if err != nil {
			return 0, err
		}

		stack.Push(res)
	}

	result, _ := stack.Pop()
	return result.(float64), nil
}

func getPriorityOperation(op rune) int {
	switch op {
	case '(':
		return 0
	case '+':
		fallthrough
	case '-':
		return 1
	case '*':
		fallthrough
	case '/':
		return 2
	default:
		return -1
	}
}

func getExpressionResult(lValue, rValue float64, opr string) (float64, error) {
	var res float64

	switch opr {
	case "+":
		res = lValue + rValue
	case "-":
		res = lValue - rValue
	case "*":
		res = lValue * rValue
	case "/":
		if rValue == 0 {
			return 0, errors.New("division by zero")
		}
		res = lValue / rValue
	}

	return res, nil
}

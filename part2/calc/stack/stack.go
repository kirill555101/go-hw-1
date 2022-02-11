package stack

import "errors"

type Stack struct {
	buff []interface{}
}

func New() *Stack {
	return &Stack{buff: make([]interface{}, 0)}
}

func (stack *Stack) Len() int {
	return len(stack.buff)
}

func (stack *Stack) Pop() (interface{}, error) {
	len := len(stack.buff)
	if len == 0 {
		return ' ', errors.New("empty stack")
	}

	popEl := stack.buff[len-1]
	stack.buff = stack.buff[:len-1]
	return popEl, nil
}

func (stack *Stack) Push(val interface{}) {
	stack.buff = append(stack.buff, val)
}

func (stack *Stack) Peek() (interface{}, error) {
	len := len(stack.buff)
	if len == 0 {
		return ' ', errors.New("empty stack")
	}

	return stack.buff[len-1], nil
}

package ch08

import (
	"errors"
	"strconv"
)

type IStack interface {
	Push(int) error
	Pop() (int, error)
	Values() []int
}

var _ IStack = (*Stack)(nil)

const maxStackSize = 100000

type Stack struct {
	values [maxStackSize]int
	top    int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Values() []int {
	return s.values[:s.top]
}

func (s *Stack) isEmpty() bool {
	return s.top == 0
}

func (s *Stack) isFull() bool {
	return s.top == maxStackSize
}

func (s *Stack) Push(x int) error {
	if s.isFull() {
		return errors.New("stack is full")
	}
	s.values[s.top] = x
	s.top++
	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("stack is empty")
	}
	s.values[s.top] = 0
	s.top--
	return s.values[s.top], nil
}

var _ IStack = (*LinkedListStack)(nil)

type LinkedListStack struct {
	values *LinkedList
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{values: NewLinkedList()}
}

func (s *LinkedListStack) Values() []int {
	size := s.values.Size()
	values := make([]int, size)
	for i := 0; i < size; i++ {
		x, _ := strconv.Atoi(s.values.GetNode(i).value)
		values[i] = x
	}
	return values
}

func (s *LinkedListStack) Push(x int) error {
	s.values.Append(NewLinkedListNode("", strconv.Itoa(x)))
	return nil
}

func (s *LinkedListStack) Pop() (int, error) {
	last := s.values.getLastNode()
	s.values.Erase(last)
	x, _ := strconv.Atoi(last.value)
	return x, nil
}

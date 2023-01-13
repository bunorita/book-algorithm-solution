package ch08

import (
	"errors"
)

const maxStackSize = 100000

type Stack struct {
	values [maxStackSize]int
	top    int
}

func NewStack(nums ...int) (*Stack, error) {
	s := &Stack{}
	for _, num := range nums {
		if err := s.Push(num); err != nil {
			return nil, err
		}
	}
	return s, nil
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

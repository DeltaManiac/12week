package queue_with_stack

/************ STACK **************/

import "errors"

type stack struct {
	s []int
}

func NewStack() *stack {
	return &stack{make([]int, 0)}
}

func (s *stack) Push(v int) {
	s.s = append(s.s, v)
}

func (s *stack) Pop() (int, error) {
	l := len(s.s)
	if l == 0 {
		return 0, errors.New("Empty Stack")
	}
	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}

func (s *stack) Len() int {
	return len(s.s)
}

/***********STACK END**************/

type MyQueue struct {
	push *stack
	pop  *stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		push: NewStack(),
		pop:  NewStack(),
	}
}

/** Push element x to the back of queue. */
func (s *MyQueue) Push(x int) {
	s.push.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (s *MyQueue) Pop() int {
	if s.pop.Len() == 0 {
		for s.push.Len() > 0 {
			v, _ := s.push.Pop()
			s.pop.Push(v)
		}
	}
	x, _ := s.pop.Pop()
	return x
}

/** Get the front element. */
func (s *MyQueue) Peek() int {
	x := s.Pop()
	s.pop.Push(x)
	return x
}

/** Returns whether the queue is empty. */
func (s *MyQueue) Empty() bool {
	if s.push.Len() == 0 && s.pop.Len() == 0 {
		return true
	}
	return false
}

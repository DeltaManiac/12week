package stack_with_queue

//-------------------------------------------------------------------------\\
//
// Since in go an array slice are powerful they can be used as an in place
// queue.
// [x:y] elements starting at index 'x' to index'y'
// [:y] all the elements except the element at 0
// [1:y] all the elements except the element at 0
//-------------------------------------------------------------------------\\

type MyStack struct {
	stack []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		stack: make([]int, 0),
	}
}

/** Push element x onto stack. */
func (s *MyStack) Push(x int) {
	s.stack = append(s.stack, x)
}

/** Removes the element on top of the stack and returns that element. */
func (s *MyStack) Pop() int {
	x := s.Top()
	s.stack = s.stack[:len(s.stack)-1]
	return x
}

/** Get the top element. */
func (s *MyStack) Top() int {
	return s.stack[len(s.stack)-1]
}

/** Returns whether the stack is empty. */
func (s *MyStack) Empty() bool {
	return len(s.stack) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

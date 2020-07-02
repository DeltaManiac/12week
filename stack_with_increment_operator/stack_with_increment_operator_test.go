package stack_with_increment_operator

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	capacity := 10
	stk := Constructor(10)
	if stk.data == nil {
		t.Errorf("Stack is nil, got: %v want: []int ", stk.data)
	}
	if cap(stk.data) != capacity {
		t.Errorf("Stack is capacity error, got: %v want: %v ", cap(stk.data), capacity)
	}
}

func TestPush(t *testing.T) {
	stk := Constructor(6)
	stk.Push(0)
	stk.Push(1)
	stk.Push(2)
	stk.Push(3)
	stk.Push(4)
	stk.Push(5)
	stk.Push(6)
	if stk.data[0] != 0 {
		t.Errorf("Stack data at index 0 is wrong , got: %d want: %d ", stk.data[0], 0)
	}
	stk.Push(6)
	if stk.data[5] != 5 {
		t.Errorf("Stack data at index 5 is wrong , got: %d want: %d ", stk.data[5], 5)
	}
}

func TestPop(t *testing.T) {
	stk := Constructor(2)
	stk.Push(0)
	stk.Push(1)
	stk.Push(2)
	y := stk.Pop()
	if y != 1 {
		t.Errorf("Stack pop is wrong , got: %d want: %d ", y, 2)
	}
	y = stk.Pop()
	if y != 0 {
		t.Errorf("Stack pop is wrong , got: %d want: %d ", y, 0)
	}
	y = stk.Pop()
	if y != -1 {
		t.Errorf("Stack pop is wrong for empty stack, got: %d want: %d ", y, -1)
	}
}

func TestIncrement(t *testing.T) {
	stk := Constructor(2)
	stk.Push(0)
	stk.Push(1)
	stk.Increment(8, 200)
	if stk.data[0] != 200 {
		t.Errorf("Stack data at index 0 is wrong , got: %d want: %d ", stk.data[0], 200)
	}
	if stk.data[1] != 201 {
		t.Errorf("Stack data at index 0 is wrong , got: %d want: %d ", stk.data[1], 201)
	}
	stk.Increment(1, 400)
	if stk.data[0] != 600 {
		t.Errorf("Stack data at index 0 is wrong , got: %d want: %d ", stk.data[0], 600)
	}
	if stk.data[1] != 201 {
		t.Errorf("Stack data at index 0 is wrong , got: %d want: %d ", stk.data[1], 201)
	}
}

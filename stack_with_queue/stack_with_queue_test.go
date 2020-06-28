package stack_with_queue

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	stk := Constructor()
	if stk.stack == nil {
		t.Errorf("Stack is nil, got: %v want: []int ", stk.stack)
	}
}

func TestPush(t *testing.T) {
	stk := Constructor()
	x := 2
	stk.Push(x)
	y := stk.stack[len(stk.stack)-1]
	if y != 2 {
		t.Errorf("Stack push failed, got: %d want: %d.", y, x)
	}
	stk.Push(990)
	stk.Push(1000)
	stk.Push(3000)

	y = stk.stack[len(stk.stack)-1]
	if y != 3000 {
		t.Errorf("Stack push failed, got: %d want: %d.", y, 3000)
	}
	y = stk.stack[len(stk.stack)-2]
	if y != 1000 {
		t.Errorf("Stack push failed, got: %d want: %d.", y, 1000)
	}
	// y = stk.Pop()
	y = stk.stack[len(stk.stack)-3]
	if y != 990 {
		t.Errorf("Stack push failed, got: %d want: %d.", y, 990)
	}

}

func TestPop(t *testing.T) {
	stk := Constructor()
	x := 2

	stk.Push(x)
	y := stk.Pop()
	if y != 2 {
		t.Errorf("Stack pop failed, got: %d want: %d.", y, x)
	}

	stk.Push(990)
	stk.Push(1000)
	stk.Push(3000)

	y = stk.Pop()
	if y != 3000 {
		t.Errorf("Stack pop failed, got: %d want: %d.", y, 3000)
	}

	y = stk.Pop()
	if y != 1000 {
		t.Errorf("Stack pop failed, got: %d want: %d.", y, 1000)
	}

	y = stk.Pop()
	if y != 990 {
		t.Errorf("Stack pop failed, got: %d want: %d.", y, 990)
	}
}

func TestTop(t *testing.T) {
	stk := Constructor()
	x := 2
	stk.Push(x)
	y := stk.Top()
	if y != 2 {
		t.Errorf("Stack Top failed, got: %d want: %d.", y, x)
	}

	stk.Pop()
	stk.Push(990)
	stk.Push(1000)
	stk.Push(3000)

	y = stk.Top()
	stk.Pop()
	if y != 3000 {
		t.Errorf("Stack top failed, got: %d want: %d.", y, 3000)
	}

	y = stk.Top()
	stk.Pop()
	if y != 1000 {
		t.Errorf("Stack top failed, got: %d want: %d.", y, 1000)
	}

	y = stk.Top()
	stk.Pop()
	if y != 990 {
		t.Errorf("Stack pop failed, got: %d want: %d.", y, 990)
	}
}

func TestEmpty(t *testing.T) {
	stk := Constructor()
	stk.Push(2)
	if stk.Empty() != false {
		t.Errorf("Stack Empty failed, got: true want: true")
	}
	stk.Pop()
	if stk.Empty() != true {
		t.Errorf("Stack Empty failed, got: false want: true")
	}
}

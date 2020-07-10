package dinner_plate_stacks

import "testing"

func TestConstructor(t *testing.T) {
	s := Constructor(3)
	if s.stacks == nil {
		t.Errorf("Stack creation failed, expected stack[] got %v", s.stacks)
	}
	if s.empty == nil {
		t.Errorf("Stack empty array failed, expected [] got %v", s.stacks)
	}
	if s.cap != 3 {
		t.Errorf("Stack capacity failed, expected cap got %v", s.cap)
	}
}

func TestPush(t *testing.T) {
	capac := 3
	s := Constructor(capac)
	s.Push(1)
	if s.stacks[0] == nil {
		t.Errorf("Stack creation failed, expected stack[0] got %v", s.stacks[0])
	}
	if cap(s.stacks[0]) != capac {
		t.Errorf("Stack creation failed, capacity expected got %v, want %v", cap(s.stacks[0]), capac)
	}
	if len(s.stacks[0]) != 1 {
		t.Errorf("Stack push failed, len expected got %v, want %v", len(s.stacks[0]), 1)
	}
	if s.stacks[0][len(s.stacks[0])-1] != 1 {
		t.Errorf("Stack top element failed, element expected got %v, want %v", s.stacks[0][len(s.stacks[0])-1], 1)
	}

	s.Push(2)
	if s.stacks[0] == nil {
		t.Errorf("Stack creation failed, expected stack[0] got %v", s.stacks[0])
	}
	if cap(s.stacks[0]) != capac {
		t.Errorf("Stack creation failed, capacity expected got %v, want %v", cap(s.stacks[0]), capac)
	}
	if len(s.stacks[0]) != 2 {
		t.Errorf("Stack push failed, len expected got %v, want %v", len(s.stacks[0]), 2)
	}
	if s.stacks[0][len(s.stacks[0])-1] != 2 {
		t.Errorf("Stack top element failed, element expected got %v, want %v", s.stacks[0][len(s.stacks[0])-1], 2)
	}
	s.Push(3)
	if s.stacks[0] == nil {
		t.Errorf("Stack creation failed, expected stack[0] got %v", s.stacks[0])
	}
	if cap(s.stacks[0]) != capac {
		t.Errorf("Stack creation failed, capacity expected got %v, want %v", cap(s.stacks[0]), capac)
	}
	if len(s.stacks[0]) != 3 {
		t.Errorf("Stack push failed, len expected got %v, want %v", len(s.stacks[0]), 3)
	}
	if s.stacks[0][len(s.stacks[0])-1] != 3 {
		t.Errorf("Stack top element failed, element expected got %v, want %v", s.stacks[0][len(s.stacks[0])-1], 3)
	}
	//NEW STACK SHOULD BE CREATED
	s.Push(1)
	if s.stacks[1] == nil {
		t.Errorf("Stack creation failed, expected stack[0] got %v", s.stacks[1])
	}
	if cap(s.stacks[1]) != capac {
		t.Errorf("Stack creation failed, capacity expected got %v, want %v", cap(s.stacks[1]), capac)
	}
	if len(s.stacks[1]) != 1 {
		t.Errorf("Stack push failed, len expected got %v, want %v", len(s.stacks[1]), 1)
	}
	if s.stacks[1][len(s.stacks[1])-1] != 1 {
		t.Errorf("Stack top element failed, element expected got %v, want %v", s.stacks[1][len(s.stacks[1])-1], 1)
	}
	s.Push(2)
	if s.stacks[1] == nil {
		t.Errorf("Stack creation failed, expected stack[0] got %v", s.stacks[1])
	}
	if cap(s.stacks[1]) != capac {
		t.Errorf("Stack creation failed, capacity expected got %v, want %v", cap(s.stacks[1]), capac)
	}
	if len(s.stacks[1]) != 2 {
		t.Errorf("Stack push failed, len expected got %v, want %v", len(s.stacks[1]), 2)
	}
	if s.stacks[1][len(s.stacks[1])-1] != 2 {
		t.Errorf("Stack top element failed, element expected got %v, want %v", s.stacks[1][len(s.stacks[1])-1], 2)
	}
	s.Push(3)
	if s.stacks[1] == nil {
		t.Errorf("Stack creation failed, expected stack[0] got %v", s.stacks[1])
	}
	if cap(s.stacks[1]) != capac {
		t.Errorf("Stack creation failed, capacity expected got %v, want %v", cap(s.stacks[1]), capac)
	}
	if len(s.stacks[1]) != 3 {
		t.Errorf("Stack push failed, len expected got %v, want %v", len(s.stacks[1]), 3)
	}
	if s.stacks[1][len(s.stacks[1])-1] != 3 {
		t.Errorf("Stack top element failed, element expected got %v, want %v", s.stacks[1][len(s.stacks[1])-1], 3)
	}
}

func TestPop(t *testing.T) {
	capac := 3
	s := Constructor(capac)
	y := s.Pop()
	if y != -1 {
		t.Errorf("Stack pop failed on empty stack, got %v, want %v", y, -1)
	}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	y = s.Pop()
	if y != 4 {
		t.Errorf("Stack pop failed,  got %v, want %v", y, 4)
	}
	y = s.Pop()
	if y != 3 {
		t.Errorf("Stack pop failed,  got %v, want %v", y, 3)
	}
	y = s.Pop()
	if y != 2 {
		t.Errorf("Stack pop failed,  got %v, want %v", y, 2)
	}
	y = s.Pop()
	if y != 1 {
		t.Errorf("Stack pop failed,  got %v, want %v", y, 1)
	}
	y = s.Pop()
	if y != -1 {
		t.Errorf("Stack pop failed on empty stack, got %v, want %v", y, -1)
	}
}

func TestPopAtStack(t *testing.T) {
	capac := 2
	s := Constructor(capac)
	y := s.PopAtStack(2)
	if y != -1 {
		t.Errorf("Stack pop failed on empty stack, got %v, want %v", y, -1)
	}
	y = s.PopAtStack(0)
	if y != -1 {
		t.Errorf("Stack pop failed on empty stack, got %v, want %v", y, -1)
	}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	y = s.PopAtStack(0)
	if y != 2 {
		t.Errorf("Stack pop failed,  got %v, want %v", y, 2)
	}
	y = s.PopAtStack(0)
	if y != 1 {
		t.Errorf("Stack pop failed,  got %v, want %v", y, 1)
	}
	y = s.PopAtStack(2)
	if y != 5 {
		t.Errorf("Stack pop failed,  got %v, want %v", y, 5)
	}
	y = s.PopAtStack(10)
	if y != -1 {
		t.Errorf("Stack pop failed,  got %v, want %v", y, -1)
	}
}

package max_stack

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	s := Constructor()
	if &s.data == nil {
		t.Errorf("Stack creation failed, expected data[] got %v", s.data)
	}
}

func TestPush(t *testing.T) {
	s := Constructor()
	s.Push(3)
	if s.data.head.data != 3 {
		t.Errorf("Stack push failed, got : %d at head,  want %d at head", s.data.head.data, 3)
	}
	if s.data.tail.data != 3 {
		t.Errorf("Stack push failed, got : %d at tail,  want %d at tail", s.data.tail.data, 3)
	}
	if s.data.tail.max.data != 3 {
		t.Errorf("Stack push failed max elem, got : %d at tail,  want %d at tail", s.data.tail.max.data, 3)
	}
	s.Push(5)
	if s.data.head.data != 3 {
		t.Errorf("Stack push failed, got : %d at head,  want %d at head", s.data.head.data, 3)
	}
	if s.data.tail.data != 5 {
		t.Errorf("Stack push failed, got : %d at tail,  want %d at head", s.data.tail.data, 5)
	}
	if s.data.tail.max.data != 5 {
		t.Errorf("Stack push failed max elem, got : %d at tail,  want %d at tail", s.data.tail.max.data, 5)
	}
	s.Push(6)
	if s.data.head.data != 3 {
		t.Errorf("Stack push failed, got : %d at head,  want %d at head", s.data.head.data, 3)
	}
	if s.data.tail.data != 6 {
		t.Errorf("Stack push failed, got : %d at tail,  want %d at head", s.data.tail.data, 6)
	}
	if s.data.tail.max.data != 6 {
		t.Errorf("Stack push failed max elem, got : %d at tail,  want %d at tail", s.data.tail.max.data, 6)
	}
	s.Push(4)
	if s.data.head.data != 3 {
		t.Errorf("Stack push failed, got : %d at head,  want %d at head", s.data.head.data, 3)
	}
	if s.data.tail.data != 4 {
		t.Errorf("Stack push failed, got : %d at tail,  want %d at head", s.data.tail.data, 4)
	}
	if s.data.tail.max.data != 6 {
		t.Errorf("Stack push failed max elem, got : %d at tail,  want %d at tail", s.data.tail.max.data, 6)
	}
}
func TestPop(t *testing.T) {
	s := Constructor()
	s.Push(3)
	s.Push(5)
	s.Push(6)
	s.Push(4)
	y := s.Pop()
	if y != 4 {
		t.Errorf("Stack pop failed , got : %d,   want %d ", y, 4)
	}
	if s.data.tail.max.data != 6 {
		t.Errorf("Stack push failed max elem, got : %d at tail,  want %d at tail", s.data.tail.max.data, 6)
	}
	y = s.Pop()
	if y != 6 {
		t.Errorf("Stack pop failed , got : %d,   want %d ", y, 6)
	}
	if s.data.tail.max.data != 5 {
		t.Errorf("Stack push failed max elem, got : %d at tail,  want %d at tail", s.data.tail.max.data, 5)
	}
	s.Push(7)
	s.Push(8)
	y = s.Pop()
	if y != 8 {
		t.Errorf("Stack pop failed , got : %d,   want %d ", y, 8)
	}
	if s.data.tail.max.data != 7 {
		t.Errorf("Stack push failed max elem, got : %d at tail,  want %d at tail", s.data.tail.max.data, 7)
	}
}

func TestTop(t *testing.T) {
	s := Constructor()
	s.Push(3)
	y := s.Top()
	if y != 3 {
		t.Errorf("Stack top failed, got %d, want %d", s.Top(), 3)
	}
	s.Push(5)
	y = s.Top()
	if y != 5 {
		t.Errorf("Stack top failed, got %d, want %d", s.Top(), 5)
	}
	s.Push(2)
	y = s.Top()
	if y != 2 {
		t.Errorf("Stack top failed, got %d, want %d", s.Top(), 2)
	}
	s.Pop()
	y = s.Top()
	if y != 5 {
		t.Errorf("Stack top failed, got %d, want %d", s.Top(), 5)
	}
	s.Pop()
	y = s.Top()
	if y != 3 {
		t.Errorf("Stack top failed, got %d, want %d", s.Top(), 3)
	}
}

func TestPeekMax(t *testing.T) {
	s := Constructor()
	s.Push(3)
	y := s.PeekMax()
	if y != 3 {
		t.Errorf("Stack PeekMax failed, got %d, want %d", s.PeekMax(), 3)
	}
	s.Push(5)
	y = s.PeekMax()
	if y != 5 {
		t.Errorf("Stack PeekMax failed, got %d, want %d", s.PeekMax(), 5)
	}
	s.Push(2)
	y = s.PeekMax()
	if y != 5 {
		t.Errorf("Stack PeekMax failed, got %d, want %d", s.PeekMax(), 5)
	}
	s.Pop()
	y = s.PeekMax()
	if y != 5 {
		t.Errorf("Stack PeekMax failed, got %d, want %d", s.PeekMax(), 5)
	}
	s.Pop()
	y = s.PeekMax()
	if y != 3 {
		t.Errorf("Stack PeekMax failed, got %d, want %d", s.PeekMax(), 3)
	}
}

func TestPopMax(t *testing.T) {
	s := Constructor()
	s.Push(8)
	s.Push(2)
	s.Push(1)
	s.Push(6)
	s.Push(7)
	s.Push(5)
	s.Push(9)
	x := s.PeekMax()

	y := s.PopMax()
	if y != x {
		t.Errorf("Stack PopMax failed, got %d, want %d", y, x)
	}

	x = s.PeekMax()
	if x != 8 {
		t.Errorf("Stack PeekMax after PopMas failed, got %d, want %d", s.PeekMax(), 8)
	}

	y = s.PopMax()
	if y != x {
		t.Errorf("Stack PopMax failed, got %d, want %d", y, x)
	}

	x = s.PeekMax()
	if x != 7 {
		t.Errorf("Stack PeekMax after PopMas failed, got %d, want %d", s.PeekMax(), 7)
	}

	y = s.PopMax()
	if y != x {
		t.Errorf("Stack PopMax failed, got %d, want %d", y, x)
	}

	x = s.PeekMax()
	if x != 6 {
		t.Errorf("Stack PeekMax after PopMas failed, got %d, want %d", s.PeekMax(), 6)
	}

	y = s.PopMax()
	if y != x {
		t.Errorf("Stack PopMax failed, got %d, want %d", y, x)
	}

	x = s.PeekMax()
	if x != 5 {
		t.Errorf("Stack PeekMax after PopMas failed, got %d, want %d", s.PeekMax(), 5)
	}

	y = s.PopMax()
	if y != x {
		t.Errorf("Stack PopMax failed, got %d, want %d", y, x)
	}

	x = s.PeekMax()
	if x != 2 {
		t.Errorf("Stack PeekMax after PopMas failed, got %d, want %d", s.PeekMax(), 2)
	}
}

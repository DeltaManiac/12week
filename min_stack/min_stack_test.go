package min_stack

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	s := Constructor()
	if s.data == nil {
		t.Errorf("Stack creation failed, expected data[] got %v", s.data)
	}
}

func TestPush(t *testing.T) {
	s := Constructor()
	s.Push(3)
	y := s.data[0]
	if y.val != 3 {
		t.Errorf("Stack Push failed for value, got %d, want %d", y.val, 3)
	}
	if y.currMin != 3 {
		t.Errorf("Stack Push failed for minimum, got %d, want %d", y.currMin, 3)
	}
	s.Push(4)
	y = s.data[1]
	if y.val != 4 {
		t.Errorf("Stack Push failed for value, got %d, want %d", y.val, 3)
	}
	if y.currMin != 3 {
		t.Errorf("Stack Push failed for minimum, got %d, want %d", y.currMin, 3)
	}
	s.Push(2)
	y = s.data[2]
	if y.val != 2 {
		t.Errorf("Stack Push failed for value, got %d, want %d", y.val, 3)
	}
	if y.currMin != 2 {
		t.Errorf("Stack Push failed for minimum, got %d, want %d", y.currMin, 3)
	}
}
func TestPop(t *testing.T) {
	s := Constructor()
	s.Push(3)
	s.Push(5)
	s.Push(2)
	s.Push(4)
	s.Push(1)
	if len(s.data) != 5 {
		t.Errorf("Stack pop failed len mismatch, got : %d,  want %d", len(s.data), 5)
	}
	s.Pop()
	if len(s.data) != 4 {
		t.Errorf("Stack pop failed len mismatch, got : %d,  want %d", len(s.data), 4)
	}
	s.Pop()
	s.Pop()
	s.Pop()
	if len(s.data) != 1 {
		t.Errorf("Stack pop failed len mismatch, got : %d,  want %d", len(s.data), 1)
	}
	s.Pop()
	if len(s.data) != 0 {
		t.Errorf("Stack pop failed len mismatch, got : %d,  want %d", len(s.data), 0)
	}
	s.Pop()
}

func TestTop(t *testing.T) {
	s := Constructor()
	s.Push(3)
	top := s.Top()
	if top != 3 {
		t.Errorf("Stack Top failed , got %d, want %d", top, 3)
	}
	s.Push(5)
	s.Push(2)
	top = s.Top()
	if top != 2 {
		t.Errorf("Stack Top failed , got %d, want %d", top, 2)
	}
	s.Pop()
	top = s.Top()
	if top != 5 {
		t.Errorf("Stack Top failed , got %d, want %d", top, 5)
	}
	s.Pop()
	s.Pop()
	s.Pop()
	top = s.Top()
	if top != -1 {
		t.Errorf("Stack Top failed , got %d, want %d", top, -1)
	}
}

func TestMin(t *testing.T) {
	s := Constructor()
	s.Push(3)
	top := s.GetMin()
	if top != 3 {
		t.Errorf("Stack Min failed , got %d, want %d", top, 3)
	}
	s.Push(5)
	s.Push(2)
	top = s.GetMin()
	if top != 2 {
		t.Errorf("Stack Top failed , got %d, want %d", top, 2)
	}
	s.Pop()
	top = s.GetMin()
	if top != 3 {
		t.Errorf("Stack Top failed , got %d, want %d", top, 3)
	}
	s.Pop()
	s.Pop()
	s.Pop()
	top = s.GetMin()
	if top != -1 {
		t.Errorf("Stack Top failed , got %d, want %d", top, -1)
	}
}

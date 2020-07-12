package queue_with_stack

import "testing"

func TestConstructor(t *testing.T) {
	s := Constructor()
	if &s == nil {
		t.Errorf("Queue creation failed, expected stack[] got %v", s)
	}
	if s.push == nil {
		t.Errorf("Queue push stack creation failed got %v", s.push)
	}
	if s.pop == nil {
		t.Errorf("Queue pop stack creation failed got %v", s.pop)
	}
}

func TestPush(t *testing.T) {
	s := Constructor()
	s.Push(1)
	if s.push.Len() != 1 && s.push.s[0] != 1 {
		t.Errorf("Queue push failed got Len %v and item at index 0 : %v, want len 1 and item 1", s.push.Len(), s.push.s[0])
	}
	s.Push(2)
	s.Push(3)
	s.Push(4)
	if s.push.Len() != 4 && s.pop.Len() != 0 {
		t.Errorf("Queue push failed got Len %v , want 4", s.push.Len())
	}
}

func TestPop(t *testing.T) {
	s := Constructor()
	s.Push(1)
	s.Push(2)
	x := s.Pop()
	if x != 1 {
		t.Errorf("Pop failed, want 1, got %v", x)
	}
	if s.pop.Len() != 1 {
		t.Errorf("Pop operation pop-stack failed, want 1, got %v", s.pop.Len())
	}
	if s.push.Len() != 0 {
		t.Errorf("Pop operation push-stack failed, want 0, got %v", s.push.Len())
	}
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Push(6)
	s.Push(7)
	if s.pop.Len() != 1 {
		t.Errorf("Pop operation pop-stack failed, want 1, got %v", s.pop.Len())
	}
	if s.push.Len() != 5 {
		t.Errorf("Pop operation push-stack failed, want 5, got %v", s.push.Len())
	}
	x = s.Pop()
	if x != 2 {
		t.Errorf("Pop failed, want 2, got %v", x)
	}
	if s.pop.Len() != 0 {
		t.Errorf("Pop operation pop-stack failed, want 0, got %v", s.pop.Len())
	}
	x = s.Pop()
	if x != 3 {
		t.Errorf("Pop failed, want 3, got %v", x)
	}
	if s.pop.Len() != 4 {
		t.Errorf("Pop operation pop-stack failed, want 4, got %v", s.pop.Len())
	}
	if s.push.Len() != 0 {
		t.Errorf("Pop operation push-stack failed, want 0, got %v", s.push.Len())
	}
	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()
	if s.pop.Len() != 0 {
		t.Errorf("Pop operation pop-stack failed, want 0, got %v", s.pop.Len())
	}
	if s.push.Len() != 0 {
		t.Errorf("Pop operation push-stack failed, want 0, got %v", s.push.Len())
	}
}

func TestPeek(t *testing.T) {
	s := Constructor()
	s.Push(1)
	s.Push(2)
	x := s.Peek()
	if x != 1 {
		t.Errorf("Peek failed, want 1, got %v", x)
	}
	s.Pop()
	x = s.Peek()
	if x != 2 {
		t.Errorf("Peek failed, want 2, got %v", x)
	}
	s.Push(3)
	x = s.Peek()
	if x != 2 {
		t.Errorf("Peek failed, want 2, got %v", x)
	}
	s.Push(4)
	s.Pop()
	x = s.Peek()
	if x != 3 {
		t.Errorf("Peek failed, want 3, got %v", x)
	}
}

func TestEmpty(t *testing.T) {
	s := Constructor()
	if s.Empty() != true {
		t.Errorf("Queue Empty failed. expected true, got %v", s.Empty())
	}
	s.Push(1)
	if s.Empty() != false {
		t.Errorf("Queue Empty failed. expected false, got %v", s.Empty())
	}

	s.Peek()
	if s.Empty() != false {
		t.Errorf("Queue Empty failed. expected false, got %v", s.Empty())
	}
	s.Pop()
	if s.Empty() != true {
		t.Errorf("Queue Empty failed. expected true, got %v", s.Empty())
	}
	s.Push(1)
	if s.Empty() != false {
		t.Errorf("Queue Empty failed. expected false, got %v", s.Empty())
	}
}

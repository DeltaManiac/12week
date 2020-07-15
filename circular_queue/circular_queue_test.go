package circular_queue

import "testing"

func TestConstructor(t *testing.T) {
	s := Constructor(4)
	if &s == nil {
		t.Errorf("Circular Queue creation failed, expected stack[] got %v", s)
	}
	if s.data == nil {
		t.Errorf("Circular Queue data creation failed got %v", s.data)
	}
	if cap(s.data) != 4 {
		t.Errorf("Circular Queue data capacity failed got %v, want 4", cap(s.data))
	}
	if s.front != 0 {
		t.Errorf("Circular Queue Queue.front failed got %v, want 0", s.front)
	}
	if s.back != -1 {
		t.Errorf("Circular Queue Queue.back failed got %v, want 0", s.back)
	}
}

func TestEnqueue(t *testing.T) {
	s := Constructor(3)
	x := s.EnQueue(1)
	if x == false {
		t.Errorf("Circular Queue  Enqueue failed got false,want true")
	}
	if s.front != 0 {
		t.Errorf("Circular Queue  Enqueue s.front failed  got %v,want %v", s.front, 0)
	}
	if s.back != 0 {
		t.Errorf("Circular Queue  Enqueue s.back failed  got %v,want %v", s.back, 0)
	}
	if s.data[0] != 1 {
		t.Errorf("Circular Queue  Enqueue s.data[%v] failed  got %v,want %v", 0, s.data[0], 1)
	}
	x = s.EnQueue(2)
	if x == false {
		t.Errorf("Circular Queue  Enqueue failed got false,want true")
	}
	if s.front != 0 {
		t.Errorf("Circular Queue  Enqueue s.front failed  got %v,want %v", s.front, 0)
	}
	if s.back != 1 {
		t.Errorf("Circular Queue  Enqueue s.back failed  got %v,want %v", s.back, 1)
	}
	if s.data[1] != 2 {
		t.Errorf("Circular Queue  Enqueue s.data[%v] failed  got %v,want %v", 1, s.data[1], 2)
	}
	x = s.EnQueue(3)
	if x == false {
		t.Errorf("Circular Queue  Enqueue failed got false,want true")
	}
	if s.front != 0 {
		t.Errorf("Circular Queue  Enqueue s.front failed  got %v,want %v", s.front, 0)
	}
	if s.back != 2 {
		t.Errorf("Circular Queue  Enqueue s.back failed  got %v,want %v", s.back, 2)
	}
	if s.data[2] != 3 {
		t.Errorf("Circular Queue  Enqueue s.data[%v] failed  got %v,want %v", 2, s.data[2], 3)
	}
	x = s.EnQueue(4)
	if x == true {
		t.Errorf("Circular Queue  Enqueue failed got true,want false")
	}
	if s.front != 0 {
		t.Errorf("Circular Queue  Enqueue s.front failed  got %v,want %v", s.front, 0)
	}
	if s.back != 2 {
		t.Errorf("Circular Queue  Enqueue s.back failed  got %v,want %v", s.back, 2)
	}
	if s.data[2] != 3 {
		t.Errorf("Circular Queue  Enqueue s.data[%v] failed  got %v,want %v", 2, s.data[2], 3)
	}
}

func TestDequeue(t *testing.T) {
	s := Constructor(2)
	x := false
	s.EnQueue(1)
	s.EnQueue(2)
	x = s.DeQueue()
	if x == false {
		t.Errorf("Circular Queue  Dequeue failed got false,want true")
	}
	if s.front != 1 {
		t.Errorf("Circular Queue  Dequeue s.front failed  got %v,want %v", s.front, 1)
	}
	if s.back != 1 {
		t.Errorf("Circular Queue  Dequeue s.back failed  got %v,want %v", s.back, 1)
	}
	s.EnQueue(3)
	if s.front != 1 {
		t.Errorf("Circular Queue  Dequeue s.front failed  got %v,want %v", s.front, 1)
	}
	if s.back != 0 {
		t.Errorf("Circular Queue  Dequeue s.back failed  got %v,want %v", s.back, 0)
	}
	x = s.DeQueue()
	if x == false {
		t.Errorf("Circular Queue  Dequeue failed got false,want true")
	}
	if s.front != 0 {
		t.Errorf("Circular Queue  Dequeue s.front failed  got %v,want %v", s.front, 0)
	}
	if s.back != 0 {
		t.Errorf("Circular Queue  Dequeue s.back failed  got %v,want %v", s.back, 0)
	}
	s.EnQueue(4)
	if s.front != 0 {
		t.Errorf("Circular Queue  Dequeue s.front failed  got %v,want %v", s.front, 0)
	}
	if s.back != 1 {
		t.Errorf("Circular Queue  Dequeue s.back failed  got %v,want %v", s.back, 1)
	}
	x = s.DeQueue()
	if x == false {
		t.Errorf("Circular Queue  Dequeue failed got false,want true")
	}
	if s.front != 1 {
		t.Errorf("Circular Queue  Dequeue s.front failed  got %v,want %v", s.front, 1)
	}
	if s.back != 1 {
		t.Errorf("Circular Queue  Dequeue s.back failed  got %v,want %v", s.back, 0)
	}
	x = s.DeQueue()
	if x == false {
		t.Errorf("Circular Queue  Dequeue failed got false,want true")
	}
	if s.front != 0 {
		t.Errorf("Circular Queue  Dequeue s.front failed  got %v,want %v", s.front, 0)
	}
	if s.back != 1 {
		t.Errorf("Circular Queue  Dequeue s.back failed  got %v,want %v", s.back, 1)
	}
	x = s.DeQueue()
	if x == true {
		t.Errorf("Circular Queue  Dequeue failed got false,want true")
	}
	if s.front != 0 {
		t.Errorf("Circular Queue  Dequeue s.front failed  got %v,want %v", s.front, 0)
	}
	if s.back != 1 {
		t.Errorf("Circular Queue  Dequeue s.back failed  got %v,want %v", s.back, 1)
	}
}

func TestRear(t *testing.T) {
	s := Constructor(2)
	x := s.Rear()
	if x != -1 {
		t.Errorf("Circular Queue  Rear failed got %v,want %v", s.Rear(), -1)
	}
	s.DeQueue()
	x = s.Rear()
	if x != -1 {
		t.Errorf("Circular Queue  Rear failed got %v,want %v", s.Rear(), -1)
	}
	s.EnQueue(1)
	x = s.Rear()
	if x != 1 {
		t.Errorf("Circular Queue  Rear failed got %v,want %v", s.Rear(), 1)
	}
	s.DeQueue()
	x = s.Rear()
	if x != -1 {
		t.Errorf("Circular Queue  Rear failed got %v,want %v", s.Rear(), -1)
	}
	s.EnQueue(0)
	s.EnQueue(1)
	x = s.Rear()
	if x != 1 {
		t.Errorf("Circular Queue  Rear failed got %v,want %v", s.Rear(), 1)
	}
	s.DeQueue()
	x = s.Rear()
	if x != 1 {
		t.Errorf("Circular Queue  Rear failed got %v,want %v", s.Rear(), 1)
	}
	s.DeQueue()
	x = s.Rear()
	if x != -1 {
		t.Errorf("Circular Queue  Rear failed got %v,want %v", s.Rear(), -1)
	}
}

func TestFront(t *testing.T) {
	s := Constructor(2)
	x := s.Front()
	if x != -1 {
		t.Errorf("Circular Queue  Front failed got %v,want %v", s.Front(), -1)
	}
	s.DeQueue()
	x = s.Front()
	if x != -1 {
		t.Errorf("Circular Queue  Front failed got %v,want %v", s.Front(), -1)
	}
	s.EnQueue(1)
	x = s.Front()
	if x != 1 {
		t.Errorf("Circular Queue  Front failed got %v,want %v", s.Front(), 1)
	}
	s.DeQueue()
	x = s.Front()
	if x != -1 {
		t.Errorf("Circular Queue  Front failed got %v,want %v", s.Front(), -1)
	}
	s.EnQueue(0)
	s.EnQueue(1)
	x = s.Front()
	if x != 0 {
		t.Errorf("Circular Queue  Front failed got %v,want %v", s.Front(), 0)
	}
	s.DeQueue()
	x = s.Front()
	if x != 1 {
		t.Errorf("Circular Queue  Front failed got %v,want %v", s.Front(), 1)
	}
	s.DeQueue()
	x = s.Front()
	if x != -1 {
		t.Errorf("Circular Queue  Front failed got %v,want %v", s.Front(), -1)
	}
}

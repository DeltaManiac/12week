package min_stack

type StackValue struct {
	val     int
	currMin int
}
type MinStack struct {
	data []StackValue
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data: make([]StackValue, 0),
	}
}

func (s *MinStack) Push(x int) {
	val := StackValue{
		val:     x,
		currMin: x,
	}
	if len(s.data) != 0 {
		t := s.top()
		if t.currMin < x {
			val.currMin = t.currMin
		}
	}
	s.data = append(s.data, val)
}

func (s *MinStack) Pop() {
	if len(s.data) > 0 {
		s.data = s.data[:len(s.data)-1]
	}
}
func (s *MinStack) top() *StackValue {
	if len(s.data) > 0 {
		return &s.data[len(s.data)-1]
	}
	return nil
}
func (s *MinStack) Top() int {
	top := s.top()
	if top != nil {
		return top.val
	}
	return -1
}

func (s *MinStack) GetMin() int {
	top := s.top()
	if top != nil {
		return top.currMin
	}
	return -1

}

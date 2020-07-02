package stack_with_increment_operator

type CustomStack struct {
	data []int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		data: make([]int, 0, maxSize),
	}
}

func (s *CustomStack) Push(x int) {
	if len(s.data) < cap(s.data) {
		s.data = append(s.data, x)
	}
}

func (s *CustomStack) Pop() int {
	if len(s.data) > 0 {
		x := s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
		return x
	}
	return -1
}

func (s *CustomStack) Increment(k int, val int) {
	count := k
	if len(s.data) < k {
		count = len(s.data)
	}
	for i := 0; i < count; i++ {
		s.data[i] = s.data[i] + val
	}
}

package dinner_plate_stacks

import "sort"

type DinnerPlates struct {
	empty  []int
	cap    int
	stacks [][]int
}

func Constructor(capacity int) DinnerPlates {
	return DinnerPlates{
		empty:  make([]int, 0),
		cap:    capacity,
		stacks: make([][]int, 0),
	}
}

func insertSort(data []int, el int) []int {
	index := sort.Search(len(data), func(i int) bool { return data[i] > el })
	data = append(data, el)
	copy(data[index+1:], data[index:])
	data[index] = el
	return data
}

func (s *DinnerPlates) Push(val int) {
	// index := s.findStack()
	if len(s.empty) > 0 && len(s.stacks[s.empty[0]]) < s.cap {
		index := s.empty[0]
		s.stacks[index] = append(s.stacks[index], val)
		if len(s.stacks[index]) == s.cap {
			s.empty = s.empty[1:]
		}
	} else {
		a := make([]int, 0, s.cap)
		a = append(a, val)
		s.stacks = append(s.stacks, a)
		if len(s.stacks[len(s.stacks)-1]) < s.cap {
			s.empty = insertSort(s.empty, len(s.stacks)-1)
		}
	}
}

func (s *DinnerPlates) findRightMostStack() int {
	// here we have to find the rightmost stack that has elements in it
	for i := len(s.stacks) - 1; i >= 0; i-- {
		if len(s.stacks[i]) > 0 {
			return i
		}
	}
	return -1
}

func (s *DinnerPlates) Pop() int {
	index := s.findRightMostStack()
	if index != -1 {
		x := s.stacks[index][len(s.stacks[index])-1]
		s.stacks[index] = s.stacks[index][:len(s.stacks[index])-1]
		s.empty = insertSort(s.empty, index)
		return x

	}
	return -1
}

func (s *DinnerPlates) PopAtStack(index int) int {
	if s.stacks == nil || index < 0 || len(s.stacks) == 0 || len(s.stacks) < index || len(s.stacks[index]) <= 0 {
		return -1
	}
	s.empty = insertSort(s.empty, index)
	x := s.stacks[index][len(s.stacks[index])-1]
	s.stacks[index] = s.stacks[index][:len(s.stacks[index])-1]
	return x
}

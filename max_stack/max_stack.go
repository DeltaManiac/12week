package max_stack

type Node struct {
	max  *Node
	data int
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
	len  int
}

func DD() DoublyLinkedList {
	return DoublyLinkedList{
		head: nil,
		tail: nil,
		len:  0,
	}
}

// Insert appends a node at the beginning of the list.
func (d *DoublyLinkedList) Insert(val int) {
	n := Node{
		data: val,
		max:  nil,
		prev: nil,
		next: nil,
	}
	n.max = &n
	if d.len == 0 {
		d.head = &n
		d.tail = &n
	} else {
		if val <= d.tail.max.data {
			n.max = d.tail.max
		}
		n.prev = d.tail
		d.tail.next = &n
		d.tail = &n
	}
	d.len++
}

func (d *DoublyLinkedList) DeleteTail() int {
	n := d.tail
	if n.prev != nil {
		n.prev.next = nil
	}
	val := d.tail.data
	d.tail = n.prev
	d.len--
	return val
}

type MaxStack struct {
	data DoublyLinkedList
}

func Constructor() MaxStack {
	return MaxStack{
		data: DD(),
	}
}

func (s *MaxStack) Push(x int) {
	s.data.Insert(x)
}

func (s *MaxStack) Pop() int {
	return s.data.DeleteTail()
}

func (s *MaxStack) Top() int {
	return s.data.tail.data
}

func (s *MaxStack) PeekMax() int {
	return s.data.tail.max.data
}

func (s *MaxStack) PopMax() int {
	//If the last one is the max return that
	if s.data.tail.max == s.data.tail {
		t := s.data.tail.data
		s.data.tail = s.data.tail.prev
		s.data.tail.next = nil
		return t
	}

	//If the first one is the max then remove and update all
	if s.data.tail.max == s.data.head {
		t := s.data.head.data
		s.data.head = s.data.head.next
		s.data.head.prev = nil
		currMaxNode := s.data.head
		s.data.head.max = currMaxNode
		ptr := s.data.head.next
		for ptr != nil {
			if ptr.data > currMaxNode.max.data {
				ptr.max = ptr
				currMaxNode = ptr
			} else {
				ptr.max = currMaxNode
			}
			ptr = ptr.next
		}
		return t
	}

	// Else its somwhere in the middle we have to do gymnastics
	// we first update all the node after the max node with the new max value
	// from the previous node and then we delete the node
	node := s.data.tail.max
	t := node.data
	maxNode := node.prev.max
	ptr := node.next
	for ptr != nil {
		if ptr.data > maxNode.data {
			ptr.max = ptr
			maxNode = ptr
		} else {
			ptr.max = maxNode
		}
		ptr = ptr.next
	}
	// Here we delete the node
	node.prev.next = node.next
	node.next.prev = node.prev
	return t
}

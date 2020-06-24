package linkedlist

/*
MyLinkedList Holds a head pointing to linked list
*/
type MyLinkedList struct {
	Head *Node
	Size int
}

/*
Node Represents an item in the linked list
*/
type Node struct {
	val  int
	next *Node
}

/*
Constructor creates an empty linked list.
*/
func Constructor() MyLinkedList {
	return MyLinkedList{
		Head: nil,
		Size: 0,
	}
}

/*
Get Retrieves the value of the index-th node in the linked list. If the index is invalid, r eturn -1.
*/
func (ll *MyLinkedList) Get(index int) int {
	if index >= ll.Size || ll.Head == nil {
		return -1

	}
	ptr := ll.Head
	for i := 0; i < index; i++ {

		ptr = ptr.next
	}
	return ptr.val
}

/*
AddAtHead Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
*/
func (ll *MyLinkedList) AddAtHead(val int) {

	node := Node{
		val:  val,
		next: nil,
	}
	if ll.Head == nil {
		ll.Head = &node
		ll.Size = 1
	} else {
		temp := ll.Head
		node.next = temp
		ll.Head = &node
		ll.Size = ll.Size + 1
	}
}

/*
AddAtTail Append a node of value val to the last element of the linked list.
*/
func (ll *MyLinkedList) AddAtTail(val int) {
	if ll.Head == nil {
		ll.AddAtHead(val)
	} else {
		ptr := ll.Head
		for ptr.next != nil {
			ptr = ptr.next
		}
		node := Node{
			val:  val,
			next: nil,
		}
		ptr.next = &node
		ll.Size = ll.Size + 1
	}

}

/*
AddAtIndex Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted.
*/
func (ll *MyLinkedList) AddAtIndex(index int, val int) {
	if index > ll.Size {
		return
	} else if index == ll.Size {
		ll.AddAtTail(val)
	} else {
		if index == 0 {
			ll.AddAtHead(val)
		} else {
			ptr := ll.Head
			i := 0
			for i < index-1 {
				ptr = ptr.next
				i++
			}
			node := Node{
				val:  val,
				next: ptr.next,
			}
			ptr.next = &node
			ll.Size = ll.Size + 1
		}
	}
}

/*
DeleteAtIndex Delete the index-th node in the linked list, if the index is valid.
*/
func (ll *MyLinkedList) DeleteAtIndex(index int) {
	if ll.Size == 0 {
		return
	}
	if index >= ll.Size {
		return
	}
	if index == 0 {
		ll.Head = ll.Head.next
		ll.Size = ll.Size - 1
		return
	}
	ptr := ll.Head
	i := 0
	for i < index-1 {
		ptr = ptr.next
		i++
	}
	if ptr.next.next == nil {
		ptr.next = nil
	} else {
		ptr.next = ptr.next.next
	}
	ll.Size = ll.Size - 1

}

//  /**
//   * Your MyLinkedList object will be instantiated and called as such:
//   * obj := Constructor();
//   * param_1 := obj.Get(index);
//   * obj.AddAtHead(val);
//   * obj.AddAtTail(val);
//   * obj.AddAtIndex(index,val);
//   * obj.DeleteAtIndex(index);
//   **/

package linkedlist

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	ll := Constructor()
	if ll.Head != nil {
		func TestConstructor(t *testing.T) {
			ll := Constructor()
			if ll.Head != nil {
				t.Errorf("LinkedList Head was incorrect, got: %v want: nil.", ll.Head)
			}
			if ll.Size != 0 {
				t.Errorf("LinkedList Size was incorrect, got: %v want: 0.", ll.Size)
			}
		}
			}
	if ll.Size != 0 {
		t.Errorf("LinkedList Size was incorrect, got: %v want: 0.", ll.Size)
	}
}

func TestAddAtHeadEmptyList(t *testing.T) {
	ll := Constructor()
	ll.AddAtHead(5)
	if ll.Head.val != 5 {
		t.Errorf("LinkedList Head was incorrect, got: %v want: %d.", ll.Head.val, 5)
	}
	if ll.Size != 1 {
		t.Errorf("LinkedList Size was incorrect, got: %v want: %d.", ll.Size, 1)
	}
}
func TestAddAtHeadNonEmptyList(t *testing.T) {
	ll := Constructor()
	ll.AddAtHead(5)
	ll.AddAtHead(6)
	if ll.Head.val != 6 {
		t.Errorf("LinkedList Head was incorrect, got: %v want: %d.", ll.Head.val, 6)
	}
	if ll.Size != 2 {
		t.Errorf("LinkedList Size was incorrect, got: %v want: %d.", ll.Size, 2)
	}
}

func TestAddAtTailEmptyList(t *testing.T) {
	ll := Constructor()
	ll.AddAtTail(5)
	if ll.Head.val != 5 {
		t.Errorf("LinkedList Tail was incorrect, got: %v want: %d.", ll.Head.val, 5)
	}
	if ll.Size != 1 {
		t.Errorf("LinkedList Size was incorrect, got: %v want: %d.", ll.Size, 1)
	}
}

func TestAddAtTailNonEmptyList(t *testing.T) {
	ll := Constructor()
	ll.AddAtTail(5)
	ll.AddAtTail(6)
	ll.AddAtTail(7)
	ptr := ll.Head
	for ptr.next != nil {
		ptr = ptr.next
	}
	if ptr.val != 7 {
		t.Errorf("LinkedList Tail was incorrect, got: %v want: %d.", ptr.val, 7)
	}
	if ll.Size != 3 {
		t.Errorf("LinkedList Size was incorrect, got: %v want: %d.", ll.Size, 3)
	}
}

func TestAddAtIndexOutOfBoundsEmptyList(t *testing.T) {
	ll := Constructor()
	ll.AddAtIndex(1, 10)
	if ll.Head != nil {
		t.Errorf("LinkedList Head was incorrect, got: %v want: nil.", ll.Head)
	}
	if ll.Size != 0 {
		t.Errorf("LinkedList Size was incorrect, got: %v want: 0.", ll.Size)
	}
}

func TestAddAtIndexBoundsEmptyList(t *testing.T) {
	ll := Constructor()
	ll.AddAtIndex(0, 10)
	if ll.Head.val != 10 {
		t.Errorf("LinkedList Head was incorrect, got: %v want: %d.", ll.Head.val, 10)
	}
	if ll.Size != 1 {
		t.Errorf("LinkedList Size was incorrect, got: %v want: %d.", ll.Size, 1)
	}
}

func TestAddAtIndexInbetweenBoundsEmptyList(t *testing.T) {
	ll := Constructor()
	ll.AddAtIndex(0, 10)
	ll.AddAtIndex(0, 20)
	ll.AddAtIndex(0, 30)
	ll.AddAtIndex(0, 40)
	ll.AddAtIndex(1, 35)
	if ll.Head.val != 40 {
		t.Errorf("LinkedList Head was incorrect, got: %v want: %d.", ll.Head.val, 40)
	}
	if ll.Size != 5 {
		t.Errorf("LinkedList Size was incorrect, got: %v want: %d.", ll.Size, 5)
	}

}

func TestAddAtIndexOutOfBoundsNonEmptyList(t *testing.T) {
	ll := Constructor()
	ll.AddAtHead(10)
	ll.AddAtIndex(10, 0)
	if ll.Head.val != 10 {
		t.Errorf("LinkedList Head was incorrect, got: %v want:%d.", ll.Head.val, 10)
	}
	if ll.Size != 1 {
		t.Errorf("LinkedList Size was incorrect, got: %v want: 1.", ll.Size)
	}
}

func TestAddAtIndexBoundsNonEmptyList(t *testing.T) {
	ll := Constructor()
	ll.AddAtHead(10)
	ll.AddAtIndex(1, 50)
	if ll.Head.val != 10 {
		t.Errorf("LinkedList Head was incorrect, got: %v want: %d.", ll.Head.val, 10)
	}
	if ll.Head.next.val != 50 {

		t.Errorf("LinkedList Head was incorrect, got: %v want: %d.", ll.Head.next.val, 50)
	}
	if ll.Size != 2 {
		t.Errorf("LinkedList Size was incorrect, got: %v want: %d.", ll.Size, 2)
	}
}

func TestAddAtIndexInbetweenBoundsNonEmptyList(t *testing.T) {
	ll := Constructor()
	ll.AddAtHead(1)
	ll.AddAtTail(6)
	ll.AddAtIndex(1, 5)
	ll.AddAtIndex(2, 3)
	ll.AddAtIndex(1, 4)
	ll.AddAtIndex(1, 2)
	if ll.Head.val != 1 {
		t.Errorf("LinkedList Head was incorrect, got: %v want: %d.", ll.Head.val, 1)
	}
	if ll.Head.next.val != 2 {
		t.Errorf("LinkedList Node was incorrect, got: %v want: %d.", ll.Head.next.val, 2)
	}
	ptr := ll.Head
	for ptr.next != nil {
		ptr = ptr.next
	}
	if ptr.val != 6 {
		t.Errorf("LinkedList Node was incorrect, got: %v want: %d.", ptr.val, 6)
	}
	if ll.Size != 6 {
		t.Errorf("LinkedList Size was incorrect, got: %v want: %d.", ll.Size, 6)
	}
}

func TestGetEmptyList(t *testing.T) {
	ll := Constructor()
	if ll.Get(0) != -1 {
		t.Errorf("LinkedList Node was incorrect, got: %v want: %d.", ll.Get(0), -1)
	}
	if ll.Get(1) != -1 {
		t.Errorf("LinkedList Node was incorrect, got: %v want: %d.", ll.Get(1), -1)
	}
}

func TestGetNonEmptyList(t *testing.T) {
	ll := Constructor()
	ll.AddAtHead(20)
	if ll.Get(0) != 20 {
		t.Errorf("LinkedList Node was incorrect, got: %v want: %d.", ll.Get(0), 20)
	}
	ll.AddAtTail(40)
	if ll.Get(1) != 40 {
		t.Errorf("LinkedList Node was incorrect, got: %v want: %d.", ll.Get(1), 40)
	}
}

func TestDeleteAtIndexEmptyList(t *testing.T) {
	ll := Constructor()
	ll.DeleteAtIndex(0)
	ll.DeleteAtIndex(10)
	if ll.Head != nil {
		t.Errorf("LinkedList Head was incorrect, got: %v want: nil.", ll.Head)
	}
	if ll.Size != 0 {
		t.Errorf("LinkedList Size was incorrect, got: %v want: 0.", ll.Size)
	}
}

func TestDeleteAtIndexNonEmptyList(t *testing.T) {
	ll := Constructor()
	ll.AddAtHead(1)
	ll.AddAtHead(2)
	ll.AddAtHead(3)
	ll.AddAtHead(4)
	ll.AddAtHead(5)
	ll.DeleteAtIndex(0)
	if ll.Get(0) != 4 {
		t.Errorf("LinkedList Node was incorrect, got: %v want: %d.", ll.Get(0), 4)
	}
	if ll.Size != 4 {

		t.Errorf("LinkedList Size was incorrect, got: %v want: 4.", ll.Size)
	}
	ll.DeleteAtIndex(2)
	if ll.Get(2) != 1 {
		t.Errorf("LinkedList Node was incorrect, got: %v want: %d.", ll.Get(2), 1)
	}
	if ll.Size != 3 {
		t.Errorf("LinkedList Size was incorrect, got: %v want: 3.", ll.Size)
	}
	ll.DeleteAtIndex(2)
	if ll.Size != 2 {
		t.Errorf("LinkedList Size was incorrect, got: %v want: 2.", ll.Size)
	}
	ll.DeleteAtIndex(12)
	if ll.Size != 2 {
		t.Errorf("LinkedList Size was incorrect, got: %v want: 2.", ll.Size)
	}
}

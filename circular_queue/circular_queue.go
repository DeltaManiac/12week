package circular_queue

type MyCircularQueue struct {
	data  []int
	front int
	back  int
	len   int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		data:  make([]int, k),
		len:   0,
		front: 0,
		back:  -1,
	}

}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}

	q.back++
	if q.back == cap(q.data) {
		q.back = 0
	}

	q.data[q.back] = value
	q.len++

	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}

	q.data[q.front] = 0

	q.front++
	if q.front == cap(q.data) {
		q.front = 0
	}

	q.len--

	return true
}

/** Get the front item from the queue. */
func (q *MyCircularQueue) Front() int {
	if !q.IsEmpty() {
		return q.data[q.front]
	}
	return -1
}

/** Get the last item from the queue. */
func (q *MyCircularQueue) Rear() int {
	if !q.IsEmpty() {
		return q.data[q.back]
	}
	return -1
}

/** Checks whether the circular queue is empty or not. */
func (q *MyCircularQueue) IsEmpty() bool {
	return q.len == 0
}

/** Checks whether the circular queue is full or not. */
func (q *MyCircularQueue) IsFull() bool {
	return q.len == cap(q.data)
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */

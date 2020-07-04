package lru_cache

type Node struct {
	data int
	key  int
	prev *Node
	next *Node
}

// Holds a doubly Linked list`
type DoublyLinkedList struct {
	head *Node
	tail *Node
	len  int
	cap  int
}

func DD(cap int) DoublyLinkedList {
	return DoublyLinkedList{
		head: nil,
		tail: nil,
		cap:  cap,
		len:  0,
	}
}

// Insert appends a node at the beginning of the list.
func (d *DoublyLinkedList) Insert(key int, val int) {
	if d.len < d.cap {
		n := Node{
			data: val,
			key:  key,
			prev: nil,
			next: nil,
		}
		if d.len == 0 {
			d.head = &n
			d.tail = &n
		} else {
			n.next = d.head
			d.head.prev = &n
			d.head = &n
		}
		d.len++
	}
}

// Delete removes the node at the end of the list which is the least recent accessed node
func (d *DoublyLinkedList) Delete() int {
	n := d.tail
	if n.prev != nil {
		n.prev.next = nil
	}
	d.tail = n.prev
	d.len--
	return n.key
}

type LRUCache struct {
	list  DoublyLinkedList
	cache map[int]*Node
	cap   int
}

func (c *LRUCache) Put(key int, value int) {
	if c.Get(key) != -1 {
		c.list.head.data = value
		return
	}
	if len(c.cache) >= c.cap {
		delete(c.cache, c.list.Delete())
	}
	c.list.Insert(key, value)
	c.cache[key] = c.list.head
}

func (c *LRUCache) Get(key int) int {
	if c.cache[key] == nil {
		return -1
	}
	if c.list.head.key == key {
		return c.list.head.data
	}
	node := c.cache[key]
	if node.next != nil {
		node.next.prev = node.prev
	}
	node.prev.next = node.next
	if c.list.tail.key == node.key {
		c.list.tail = node.prev
	}
	c.list.len--
	c.list.Insert(node.key, node.data)
	c.cache[key] = c.list.head
	return node.data
}

func Constructor(cap int) LRUCache {
	return LRUCache{
		list:  DD(cap),
		cache: make(map[int]*Node),
		cap:   cap,
	}
}

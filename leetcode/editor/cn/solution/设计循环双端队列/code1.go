package 设计循环双端队列

// 这是链表解法，严格上说，链表解法不对
// 因为这里循环队列，特指是在数组做队列时容易溢出，数组元素不满所设计的
// 专门用数组来解得算法

// 这可以看做是双端队列的一种实现
type Deque struct {
	cap int
	len int

	head *DequeNode
	tail *DequeNode
}

type DequeNode struct {
	data int
	next *DequeNode
	pre  *DequeNode
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func ConstructorL(k int) Deque {
	deque := Deque{}
	deque.cap = k
	return deque
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *Deque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	p := &DequeNode{data: value}
	p.next = this.head
	if this.head != nil {
		this.head.pre = p
	}
	this.head = p
	this.len++
	if this.len == 1 {
		this.tail = p
	}
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *Deque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}

	p := &DequeNode{data: value}

	p.pre = this.tail
	if this.tail != nil {
		this.tail.next = p
	}
	this.tail = p
	this.len++
	if this.len == 1 {
		this.head = p
	}
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *Deque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}

	this.head = this.head.next
	if this.head != nil {
		this.head.pre = nil
	}
	this.len--
	if this.len == 0 {
		this.tail = nil
	}
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *Deque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}

	this.tail = this.tail.pre
	if this.tail != nil {
		this.tail.next = nil
	}
	this.len--
	if this.len == 0 {
		this.head = nil
	}
	return true
}

/** Get the front item from the deque. */
func (this *Deque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.head.data
}

/** Get the last item from the deque. */
func (this *Deque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.tail.data
}

/** Checks whether the circular deque is empty or not. */
func (this *Deque) IsEmpty() bool {
	return this.len == 0
}

/** Checks whether the circular deque is full or not. */
func (this *Deque) IsFull() bool {
	return this.len >= this.cap
}

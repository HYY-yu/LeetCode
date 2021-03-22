package design_circular_deque

// 数组解法
// 关键理清四个操作
// 循环指针逆时针转：(head -1 + cap) % cap
// 循环指针顺时针转：(tail + 1) % cap
// 判空：head == tail
// 判满：(tail)
type MyCircularDeque struct {
	cap  int
	head int // 逆时针转
	tail int // 顺时针转

	data []int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) *MyCircularDeque {
	deque := MyCircularDeque{}
	deque.cap = k + 1

	deque.data = make([]int, deque.cap)
	return &deque
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}

	this.head = (this.head - 1 + this.cap) % this.cap
	this.data[this.head] = value
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}

	this.data[this.tail] = value
	this.tail = (this.tail + 1) % this.cap
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head + 1) % this.cap
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.tail = (this.tail - 1 + this.cap) % this.cap
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.head]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[(this.tail-1+this.cap)%this.cap]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.head == this.tail
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return (this.tail+1)%this.cap == this.head
}

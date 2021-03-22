// 设计实现双端队列。
// 你的实现需要支持以下操作：
//
//
// MyCircularDeque(k)：构造函数,双端队列的大小为k。
// insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true。
// insertLast()：将一个元素添加到双端队列尾部。如果操作成功返回 true。
// deleteFront()：从双端队列头部删除一个元素。 如果操作成功返回 true。
// deleteLast()：从双端队列尾部删除一个元素。如果操作成功返回 true。
// getFront()：从双端队列头部获得一个元素。如果双端队列为空，返回 -1。
// getRear()：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1。
// isEmpty()：检查双端队列是否为空。
// isFull()：检查双端队列是否满了。
//
//
// 示例：
//
// MyCircularDeque circularDeque = new MycircularDeque(3); // 设置容量大小为3
// circularDeque.insertLast(1);			        // 返回 true
// circularDeque.insertLast(2);			        // 返回 true
// circularDeque.insertFront(3);			        // 返回 true
// circularDeque.insertFront(4);			        // 已经满了，返回 false
// circularDeque.getRear();  				// 返回 2
// circularDeque.isFull();				        // 返回 true
// circularDeque.deleteLast();			        // 返回 true
// circularDeque.insertFront(4);			        // 返回 true
// circularDeque.getFront();				// 返回 4
//
//
//
//
// 提示：
//
//
// 所有值的范围为 [1, 1000]
// 操作次数的范围为 [1, 1000]
// 请不要使用内置的双端队列库。
//
// Related Topics 设计 队列
// 👍 76 👎 0

package main

import "fmt"

func main() {
	obj := Constructor(3)

	fmt.Println(obj.InsertFront(3))
	fmt.Println(obj.GetFront())
	fmt.Println(obj.GetRear())

	fmt.Println(obj.InsertLast(2))
	fmt.Println(obj.GetFront())
	fmt.Println(obj.GetRear())

	fmt.Println(obj.InsertFront(6))
	fmt.Println(obj.GetFront())
	fmt.Println(obj.GetRear())

	fmt.Println(obj.DeleteFront())
	fmt.Println(obj.GetFront())
	fmt.Println(obj.GetRear())

	fmt.Println(obj.DeleteFront())
	fmt.Println(obj.GetFront())
	fmt.Println(obj.GetRear())

	fmt.Println(obj.DeleteFront())
	fmt.Println(obj.GetFront())
	fmt.Println(obj.GetRear())

	fmt.Println(obj.DeleteFront())
	fmt.Println(obj.GetFront())
	fmt.Println(obj.GetRear())
}

// leetcode submit region begin(Prohibit modification and deletion)

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
func Constructor(k int) MyCircularDeque {
	deque := MyCircularDeque{}
	deque.cap = k + 1

	deque.data = make([]int, deque.cap)
	return deque
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

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
// leetcode submit region end(Prohibit modification and deletion)

package min_stack

import "math"

// 2. 思路：不需要辅助栈，我直接把上个最小值入栈
// 9 4 4 6 7
// min = 4
// push(0)
// 9 4 4 6 7 4 0 // 注意，上个最小值4入栈了
// min = 0
// push(0)
// 9 4 6 7 4 0 0
// min = 0
// pop 操作
// if min == pop() then min = pop() （弹出两次是关键）
type MinStackOM struct {
	stack []int
	min   int
}

/** initialize your data structure here. */
func Constructor() MinStackOM {
	m := MinStackOM{}
	m.stack = make([]int, 0)
	m.min = math.MaxInt64
	return m
}

func (this *MinStackOM) Push(x int) {
	if this.min >= x {
		this.stack = append(this.stack, this.min)
		this.min = x
	}
	this.stack = append(this.stack, x)
}

func (this *MinStackOM) Pop() {
	top := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if top == this.min {
		this.min = this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
	}
}

func (this *MinStackOM) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStackOM) GetMin() int {
	return this.min
}

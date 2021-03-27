package 最小栈

// 了解即可，无需掌握

// 3. 思路：
// 栈内存的是与最小值的差值

// push(3) push(6)
// -2 0 -3
// 差值stack : 0   2   -1
// min :   -2      -3

// push() : elem - min
// pop() : min + stackTop
type MinStackOC struct {
	stack []int
	min   int
}

/** initialize your data structure here. */
func ConstructorOC() MinStackOC {
	m := MinStackOC{}
	m.stack = make([]int, 0)
	return m
}

func (this *MinStackOC) Push(x int) {
	if len(this.stack) == 0 {
		this.min = x
	}
	this.stack = append(this.stack, x-this.min)
	if this.min > x {
		this.min = x
	}
}

func (this *MinStackOC) Pop() {
	top := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if top < 0 {
		this.min -= top
	}
}

func (this *MinStackOC) Top() int {
	top := this.stack[len(this.stack)-1]
	if top < 0 {
		return this.min
	}

	return top + this.min
}

func (this *MinStackOC) GetMin() int {
	return this.min
}

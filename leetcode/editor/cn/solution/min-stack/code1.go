package min_stack

// 1. 思路：辅助栈，与栈本身不同步
// 9 4 6 7 0 1 0
// 9 4     0 0
// 这里一定要存两个0，弹出时才不会错
// 4后面可以不存
type MinStackDouble struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func ConstructorDouble() MinStackDouble {
	m := MinStackDouble{}
	m.stack = make([]int, 0)
	m.minStack = make([]int, 0)
	return m
}

func (this *MinStackDouble) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, x)
	} else {
		minTop := this.minStack[len(this.minStack)-1]
		if minTop >= x {
			this.minStack = append(this.minStack, x)
		}
	}
}

func (this *MinStackDouble) Pop() {
	top := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if top == this.minStack[len(this.minStack)-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

func (this *MinStackDouble) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStackDouble) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

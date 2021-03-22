// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
//
//
// push(x) —— 将元素 x 推入栈中。
// pop() —— 删除栈顶的元素。
// top() —— 获取栈顶元素。
// getMin() —— 检索栈中的最小元素。
//
//
//
//
// 示例:
//
// 输入：
// ["MinStackOM","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[0],[-3],[],[],[],[]]
//
// 输出：
// [null,null,null,null,-3,null,0,-2]
//
// 解释：
// MinStackOM minStack = new MinStackOM();
// minStack.push(-2);
// minStack.push(0);
// minStack.push(-3);
// minStack.getMin();   --> 返回 -3.
// minStack.pop();
// minStack.top();      --> 返回 0.
// minStack.getMin();   --> 返回 -2.
//
//
//
//
// 提示：
//
//
// pop、top 和 getMin 操作总是在 非空栈 上调用。
//
// Related Topics 栈 设计
// 👍 835 👎 0

package main

import (
	"fmt"
)

func main() {
	minStack := ConstructorDouble()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	x := minStack.GetMin()
	fmt.Println(x)
	minStack.Pop()
	x = minStack.Top()
	fmt.Println(x)
	x = minStack.GetMin()
	fmt.Println(x)
}

// leetcode submit region begin(Prohibit modification and deletion)
// 1. 思路：辅助栈，与栈本身不同步
// 9 4 6 7 0 1 0
// 9 4     0 0
// 这里一定要存两个0
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

/**
 * Your MinStackOM object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// leetcode submit region end(Prohibit modification and deletion)

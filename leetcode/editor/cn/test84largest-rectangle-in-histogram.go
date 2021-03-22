// 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。
//
//
//
//
//
// 以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。
//
//
//
//
//
// 图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。
//
//
//
// 示例:
//
// 输入: [2,1,5,6,2,3]
// 输出: 10
// Related Topics 栈 数组
// 👍 1235 👎 0

package main

import "fmt"

func main() {
	a := []int{2, 1, 5, 6, 2, 3}
	r := largestRectangleArea(a)
	fmt.Println(r)
}

// leetcode submit region begin(Prohibit modification and deletion)
// 栈+哨兵
func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	// [2,1,5,6,2,3]
	// order stack ： -1 2
	// 出栈元素:
	// 其左边界 left 和 右边界 right 中间夹住的位置，就是它的最大面积

	stack := make([]int, 0, len(heights))

	heights = append([]int{-1}, heights...)
	heights = append(heights, -1) // 尾哨兵
	maxArea := 0

	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[peek(stack)] > heights[i] {
			s, p := pop(stack)
			stack = s

			if area := (i - peek(stack) - 1) * heights[p]; area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, i)
	}
	return maxArea
}

func peek(stack []int) int {
	return stack[len(stack)-1]
}

func pop(stack []int) ([]int, int) {
	p := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return stack, p
}

// leetcode submit region end(Prohibit modification and deletion)

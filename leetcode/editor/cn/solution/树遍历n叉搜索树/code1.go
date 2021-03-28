package 树遍历n叉搜索树

import (
	"container/list"
)

// N叉树
type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	result := make([]int, 0)
	return preOrderWithStack(root, result)
}

// 递归遍历
func postOrder(p *Node, result []int) []int {
	if p != nil {
		// 前序 先访问再遍历，后序先遍历再访问
		result = append(result, p.Val)
		for _, child := range p.Children {
			result = postOrder(child, result)
		}
	}
	return result
}

// 用栈遍历N叉树
// DFS 深度优先搜索
func preOrderWithStack(p *Node, result []int) []int {
	if p == nil {
		return result
	}
	stack := make([]*Node, 0)
	stack = append(stack, p)

	for len(stack) > 0 {
		stack, p = popN(stack)
		// 1. 实际访问节点代码，若是后续遍历就在弹出栈时即可
		result = append(result, p.Val)

		// 子节点入栈，注意: 从右到左
		for i := len(p.Children) - 1; i >= 0; i-- {
			stack = append(stack, p.Children[i])
		}
	}
	return result
}

// 层序遍历N叉树
// BFS 广度优先遍历
func levelOrder(root *Node) [][]int {
	result := make([][]int, 0)
	isVisit := make(map[*Node]int)

	p := root
	if p == nil {
		return result
	}
	queue := list.New()
	queue.PushBack(p)

	for queue.Len() != 0 {
		r := make([]int, 0)
		// 处理当前层，这里比较特殊，因为我们需要二维数组的结果
		// 如果只是一维结果，是不需要这个循环的
		size := queue.Len()
		for i := 0; i < size; i++ {
			e := queue.Front()
			p = e.Value.(*Node)
			queue.Remove(e)

			if _, ok := isVisit[p]; !ok {
				r = append(r, p.Val)
				isVisit[p] = 1
			}
			for _, child := range p.Children {
				queue.PushBack(child)
			}
		}
		// -- 处理完毕
		result = append(result, r)
	}
	return result
}

func popN(s []*Node) ([]*Node, *Node) {
	if len(s) == 0 {
		return s, nil
	}
	v := s[len(s)-1]
	s = s[:len(s)-1]
	return s, v
}

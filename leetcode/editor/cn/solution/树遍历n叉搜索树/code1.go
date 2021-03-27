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

func isVisited(vi []*Node, v *Node) (result bool) {
	for _, vv := range vi {
		if vv == v {
			return true
		}
	}
	return
}

// 用栈遍历N叉树
// DFS 深度优先搜索
func preOrderWithStack(p *Node, result []int) []int {
	if p == nil {
		return result
	}

	stack := make([]*Node, 0)
	stack = append(stack, p)
	visited := make([]*Node, 0)

	for p != nil || len(stack) > 0 {
		p = peekN(stack)
		// 实际访问节点代码，若是后续遍历就在弹出栈时即可
		if !isVisited(visited, p) {
			visited = append(visited, p)
			result = append(result, p.Val)
		}
		// 找个没被访问的子节点
		flag := false
		for _, child := range p.Children {
			if !isVisited(visited, child) {
				stack = append(stack, child)
				p = child
				flag = true
				break
			}
		}
		if !flag {
			// 没有没被访问的子节点（或者没有子节点），这个节点就可以弹出了
			stack, p = popN(stack)
			p = peekN(stack)
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
		// 处理当前层
		size := queue.Len()
		for i := 0; i < size; i++ {
			e := queue.Front()
			p = e.Value.(*Node)
			if _, ok := isVisit[p]; !ok {
				r = append(r, p.Val)
				isVisit[p] = 1
			}
			for _, child := range p.Children {
				queue.PushBack(child)
			}
			queue.Remove(e)
		}
		// -- 处理完毕
		result = append(result, r)
	}
	return result
}

func peekN(s []*Node) *Node {
	if len(s) == 0 {
		return nil
	}
	return s[len(s)-1]
}

func popN(s []*Node) ([]*Node, *Node) {
	if len(s) == 0 {
		return s, nil
	}
	v := s[len(s)-1]
	s = s[:len(s)-1]
	return s, v
}

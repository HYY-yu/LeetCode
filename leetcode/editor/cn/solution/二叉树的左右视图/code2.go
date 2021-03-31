package 二叉树的左右视图

// DFS
type StackNode struct {
	T     *TreeNode
	Level int
}

func rightSideView(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	stack := make([]StackNode, 0)
	stack = append(stack, StackNode{T: root, Level: 0})
	var p StackNode

	for len(stack) > 0 {
		stack, p = popS(stack)

		// 这个很难想
		if p.Level == len(result) {
			result = append(result, p.T.Val)
		}

		if p.T.Left != nil {
			stack = append(stack, StackNode{T: p.T.Left, Level: p.Level + 1})
		}
		if p.T.Right != nil {
			stack = append(stack, StackNode{T: p.T.Right, Level: p.Level + 1})
		}
	}
	return result
}

func popS(stack []StackNode) ([]StackNode, StackNode) {
	n := len(stack)
	p := stack[n-1]
	stack = stack[:n-1]
	return stack, p
}

package binary_tree_traversal

// 二叉树的前序 中序遍历
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	if root != nil {
		result = inorder(root, result)
	}
	return result
}

func inorder(p *TreeNode, r []int) []int {
	if p != nil {
		r = inorder(p.Left, r)
		r = append(r, p.Val)
		r = inorder(p.Right, r)
	}
	return r
}

func inorderWithStack(p *TreeNode, r []int) []int {
	stack := make([]*TreeNode, 0)
	for p != nil || len(stack) > 0 {
		if p == nil {
			break
		}
		r = append(r, p.Val)
		for p.Left != nil {
			stack = append(stack, p)
			p = p.Left
			r = append(r, p.Val)
		}
		stack, p = pop(stack)

		p = p.Right
	}
	return r
}

func preorderWithStack(p *TreeNode, r []int) []int {
	stack := make([]*TreeNode, 0)
	for p != nil || len(stack) > 0 {
		for p != nil {
			r = append(r, p.Val)
			stack = append(stack, p)
			p = p.Left
		}
		stack, p = pop(stack)
		p = p.Right
	}
	return r
}

func postorderWithStack(p *TreeNode, r []int) []int {
	// 等待实现
	return nil
}

func pop(s []*TreeNode) ([]*TreeNode, *TreeNode) {
	if len(s) == 0 {
		return s, nil
	}
	v := s[len(s)-1]
	s = s[:len(s)-1]
	return s, v
}

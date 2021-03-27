package 树遍历2叉搜索树

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

// 中序遍历
// 前序遍历同中序遍历，差别只在访问代码的位置
func inorderWithStack(p *TreeNode, r []int) []int {
	stack := make([]*TreeNode, 0)

	for p != nil || len(stack) > 0 {
		// 访问p的左子树直到nil
		for p != nil {
			// 前序遍历，在这里访问
			// r = append(r, p.Val)
			stack = append(stack, p)
			p = p.Left
		}
		// 弹出栈顶元素
		stack, p = pop(stack)
		//访问代码
		r = append(r, p.Val)
		// 进入右子树
		p = p.Right
	}

	return r
}

// 后序遍历
func postorderWithStack(p *TreeNode, r []int) []int {
	stack := make([]*TreeNode, 0)
	var lr *TreeNode

	for p != nil || len(stack) > 0 {
		// 当p不空 且 p不是上个弹出的元素 且 p不是栈顶元素（栈顶元素的左节点一定被访问过了）
		for p != nil && p != lr && p != peek(stack) {
			stack = append(stack, p)
			p = p.Left
		}
		p = peek(stack)

		// 当p的右子树空，或者p的右子树正好是上次弹出的元素，说明p可以被访问和弹出了（左右子树都遍历过了）
		if p.Right == nil || p.Right == lr {
			r = append(r, p.Val)
			stack, p = pop(stack)
			lr = p

			p = peek(stack)
		} else {
			p = p.Right
		}
	}
	return r
}

func peek(s []*TreeNode) *TreeNode {
	if len(s) == 0 {
		return nil
	}
	return s[len(s)-1]
}

func pop(s []*TreeNode) ([]*TreeNode, *TreeNode) {
	if len(s) == 0 {
		return s, nil
	}
	v := s[len(s)-1]
	s = s[:len(s)-1]
	return s, v
}

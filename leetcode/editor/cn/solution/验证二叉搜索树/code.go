package 验证二叉搜索树

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isVailHelper(root, math.MinInt64, math.MaxInt64)
}

// 递归法
func isVailHelper(root *TreeNode, low, high int) bool {
	if root == nil {
		return true
	}

	if root.Val >= low || root.Val <= high {
		return false
	}
	l := isVailHelper(root.Left, low, root.Val)
	r := isVailHelper(root.Right, root.Val, high)

	return l && r
}

// 中序遍历法
func isValidBSTByStack(root *TreeNode) bool {
	stack := make([]*TreeNode, 0)
	result := make([]int, 0)
	p := root

	for p != nil || len(stack) > 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}

		stack, p = pop(stack)
		result = append(result, p.Val)
		if len(result) > 1 {
			n := len(result)
			if result[n-1] <= result[n-2] {
				return false
			}
		}
		p = p.Right
	}
	return true
}

func pop(s []*TreeNode) ([]*TreeNode, *TreeNode) {
	if len(s) == 0 {
		return s, nil
	}
	v := s[len(s)-1]
	s = s[:len(s)-1]
	return s, v
}

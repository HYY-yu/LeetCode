package 二叉树的左右视图

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFS
func rightSideViewBFS(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	var p *TreeNode

	for len(queue) > 0 {
		var r int
		size := len(queue)
		for i := 0; i < size; i++ {
			queue, p = front(queue)
			if i == 0 {
				r = p.Val
			}

			if p.Left != nil {
				queue = append(queue, p.Left)
			}
			if p.Right != nil {
				queue = append(queue, p.Right)
			}
		}
		result = append(result, r)
	}
	return result
}

func front(queue []*TreeNode) ([]*TreeNode, *TreeNode) {
	if len(queue) == 0 {
		return queue, nil
	}
	p := queue[0]
	queue = queue[1:]
	return queue, p
}

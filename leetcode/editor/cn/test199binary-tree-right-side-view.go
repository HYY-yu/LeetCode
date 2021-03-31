// 给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
//
// 示例:
//
// 输入: [1,2,3,null,5,null,4]
// 输出: [1, 3, 4]
// 解释:
//
//   1            <---
// /   \
// 2     3         <---
// \     \
//  5     4       <---
//
// Related Topics 树 深度优先搜索 广度优先搜索 递归 队列
// 👍 436 👎 0

package main

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
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

// leetcode submit region end(Prohibit modification and deletion)

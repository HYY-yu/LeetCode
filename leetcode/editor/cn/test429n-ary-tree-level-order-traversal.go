//给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。
//
// 树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。
//
//
//
// 示例 1：
//
//
//
//
//输入：root = [1,null,3,2,4,null,5,6]
//输出：[[1],[3,2,4],[5,6]]
//
//
// 示例 2：
//
//
//
//
//输入：root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,
//null,13,null,null,14]
//输出：[[1],[2,3,4,5],[6,7,8,9,10],[11,12,13],[14]]
//
//
//
//
// 提示：
//
//
// 树的高度不会超过 1000
// 树的节点总数在 [0, 10^4] 之间
//
// Related Topics 树 广度优先搜索
// 👍 142 👎 0

package main

import "container/list"

func main() {
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

// N叉树
type Node struct {
	Val      int
	Children []*Node
}

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

//leetcode submit region end(Prohibit modification and deletion)

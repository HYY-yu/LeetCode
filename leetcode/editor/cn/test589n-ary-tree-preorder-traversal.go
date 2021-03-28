// 给定一个 N 叉树，返回其节点值的 前序遍历 。
//
// N 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。
//
//
//
//
//
// 进阶：
//
// 递归法很简单，你可以使用迭代法完成此题吗?
//
//
//
// 示例 1：
//
//
//
//
// 输入：root = [1,null,3,2,4,null,5,6]
// 输出：[1,3,5,6,2,4]
//
// 示例 2：
//
//
//
//
// 输入：root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,
// null,13,null,null,14]
// 输出：[1,2,3,6,7,11,14,4,8,12,5,9,13,10]
//
//
//
//
// 提示：
//
//
// N 叉树的高度小于或等于 1000
// 节点总数在范围 [0, 10^4] 内
//
//
//
// Related Topics 树
// 👍 150 👎 0

package main

func main() {
}

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	result := make([]int, 0)
	result = preOrderWithStack(root, result)
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

func popN(s []*Node) ([]*Node, *Node) {
	if len(s) == 0 {
		return s, nil
	}
	v := s[len(s)-1]
	s = s[:len(s)-1]
	return s, v
}

// leetcode submit region end(Prohibit modification and deletion)

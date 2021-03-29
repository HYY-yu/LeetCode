// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//
//
// 示例 1：
//
//
// 输入：l1 = [1,2,4], l2 = [1,3,4]
// 输出：[1,1,2,3,4,4]
//
//
// 示例 2：
//
//
// 输入：l1 = [], l2 = []
// 输出：[]
//
//
// 示例 3：
//
//
// 输入：l1 = [], l2 = [0]
// 输出：[0]
//
//
//
//
// 提示：
//
//
// 两个链表的节点数目范围是 [0, 50]
// -100 <= Node.val <= 100
// l1 和 l2 均按 非递减顺序 排列
//
// Related Topics 递归 链表
// 👍 1586 👎 0

package main

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 比较 p1 指向节点 与 p2 指向的节点
// 注意：最后记得把p1或者p2剩余节点加入链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head = &ListNode{}
	p := head
	p1, p2 := l1, l2

	for p1 != nil && p2 != nil {
		if p1.Val >= p2.Val {
			p.Next = p2
			p2 = p2.Next
		} else {
			p.Next = p1
			p1 = p1.Next
		}
		p = p.Next
	}

	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	return head.Next
}

// leetcode submit region end(Prohibit modification and deletion)

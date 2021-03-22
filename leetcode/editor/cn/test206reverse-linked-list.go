// 反转一个单链表。
//
// 示例:
//
// 输入: 1->2->3->4->5->NULL
// 输出: 5<-4->3->2->1->NULL
// 		   pr  p
// 进阶:
// 你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
// Related Topics 链表
// 👍 1554 👎 0

package main

func main() {
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var prep *ListNode
	p := head

	for p != nil {
		p.Next, prep, p = prep, p, p.Next
	}
	return prep
}

// leetcode submit region end(Prohibit modification and deletion)

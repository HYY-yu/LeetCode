package 合并有序链表

type ListNode struct {
	Val  int
	Next *ListNode
}

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

package 翻转链表

type ListNode struct {
	Val  int
	Next *ListNode
}

// 简洁明了
func reverseList(head *ListNode) *ListNode {
	var prep *ListNode
	p := head

	for p != nil {
		p.Next, prep, p = prep, p, p.Next
	}
	return prep
}

// 详细版
func reverseLinkList2(head *ListNode) *ListNode {
	var prep *ListNode
	p := head
	for p != nil {
		// 1- > 3 -> 5
		// 1 <- 3 <- 5
		p.Next = prep
		prep = p
		p = p.Next
	}
	return p
}

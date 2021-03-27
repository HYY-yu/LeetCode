package 两两交换链表节点

type ListNode struct {
	Val  int
	Next *ListNode
}

// dummyHead 可以省很多空值判断，且在这题中，一个虚拟的头结点使代码更简洁
// 相当于哨兵节点
func swapPairs(head *ListNode) *ListNode {
	dummyHead := new(ListNode)
	dummyHead.Next = head

	// 就是p开始，不断交换下两个节点的值
	p := dummyHead
	for p.Next != nil && p.Next.Next != nil {
		n1 := p.Next
		n2 := p.Next.Next
		// p -> n1 -> n2 -> next
		// p -> n2 -> n1 -> next
		p.Next, n1.Next, n2.Next = n2, n2.Next, n1
		p = n1
	}
	return dummyHead.Next
}

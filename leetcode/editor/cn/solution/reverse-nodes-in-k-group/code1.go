package reverse_nodes_in_k_group

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	// 思路：
	// s p 作为 k个子链表的头尾
	// 1.逆置 s p 形成的子链表
	// 2. 循环直到 s p 的距离不满足k
	dummyHead := &ListNode{Next: head}
	var pre = dummyHead
	var p = head

	for p != nil {
		s := p
		for i := 0; i < k-1; i++ {
			p = p.Next
			if p == nil {
				return dummyHead.Next
			}
		}
		// dummy -> 1 -> 2 -> 3 -> 4 -> 5
		// pre      s    p
		// dummy -> 2 -> 1 -> 3 -> 4 -> 5
		// pre      f    e

		pn := p.Next
		f, e := reverseSingle(s, p)
		pre.Next = f
		e.Next = pn
		pre = e
		p = pn
	}
	return dummyHead.Next
}

func reverseSingle(s, e *ListNode) (*ListNode, *ListNode) {
	e.Next = nil
	// s  ->  x -> x -> e ->nil
	// pre    p

	var pre *ListNode
	p := s

	for p != nil {
		p.Next, pre, p = pre, p, p.Next
	}
	return pre, s
}

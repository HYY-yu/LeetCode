package swap_nodes_in_pairs

// 自己想的一种解法
// 没有dummyHead，比较复杂
// 就是 l r 交换，然后下个l r 因为cFlag原因跳过
// 不太优雅，不用学习
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	l, r := head, head.Next
	var ll *ListNode
	cFlag := false
	for r != nil {
		if !cFlag {
			if ll == nil {
				// l- > r- > rr
				// r- > l- > rr
				head, l.Next, r.Next = r, r.Next, l
			} else {
				// ll- > l- > r- > rr
				// ll- > r- > l- > rr
				ll.Next, l.Next, r.Next = r, r.Next, l
			}
			l, r = r, l
			cFlag = true
		} else {
			cFlag = false
		}
		ll, l, r = l, r, r.Next
	}
	return head
}

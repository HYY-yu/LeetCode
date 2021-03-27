package 判断环形链表的环入口

type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路还是快慢指针
// 当快指针遇上满指针时，说明有环，为了找到环:
// 则 slow 继续在环中运行
// fast 从 head 重新开始运行
// 当fast == flow时，必然是链表的入口处
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for {
		//
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}
	// 若链表有环
	// 则 slow 在环中运行
	// fast 从 head 重新开始运行
	// 当fast == flow时，必然是链表的入口处
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

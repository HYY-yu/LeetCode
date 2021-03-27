package 判断环形链表

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针即可
// (当快指针和慢指针重合，代表有环)
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
}

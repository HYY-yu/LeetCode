package 判断环形链表

func main() {

}

type LinkNode struct {
	Val  int
	next *LinkNode
}

func check(root *LinkNode) bool {
	var p1, p2 *LinkNode
	p1 = root
	p2 = root

	for p2 != nil && p2.next != nil {
		p1, p2 = p1.next, p2.next.next

		if p1 == p2 {
			return true
		}
	}
	return false
}

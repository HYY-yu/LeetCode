package valid_parentheses

import "container/list"

// 这里有个小考虑就是
// 如果用go的数组做栈，假设一种极端情况，就是这个s的长度很长数组需要不断扩容，创建新数组
// 如果用链表做栈则无此问题
// 链表解法
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := list.New()
	isEq := func(a, b byte) bool {
		return (a == '(' && b == ')') || (a == '[' && b == ']') || (a == '{' && b == '}')
	}

	for i := 0; i < len(s); i++ {
		if v := stack.Front(); v != nil {
			vbyte := v.Value.(byte)
			if isEq(vbyte, s[i]) {
				stack.Remove(v)
				continue
			}
		}
		stack.PushFront(s[i])
	}
	return stack.Len() == 0
}

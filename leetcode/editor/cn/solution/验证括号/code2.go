package 验证括号

// 数组做栈
func isValidArray(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]byte, 0)
	isEq := func(a, b byte) bool {
		return (a == '(' && b == ')') || (a == '[' && b == ']') || (a == '{' && b == '}')
	}

	for i := 0; i < len(s); i++ {
		if len(stack) > 0 {
			vbyte := stack[len(stack)-1]
			if isEq(vbyte, s[i]) { // 关键代码
				stack = stack[:len(stack)-1]
				continue
			}
		}
		stack = append(stack, s[i])
	}
	return len(stack) == 0
}

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
//
//
//
//
// 示例 1：
//
//
// 输入：s = "()"
// 输出：true
//
//
// 示例 2：
//
//
// 输入：s = "()[]{}"
// 输出：true
//
//
// 示例 3：
//
//
// 输入：s = "(]"
// 输出：false
//
//
// 示例 4：
//
//
// 输入：s = "([)]"
// 输出：false
//
//
// 示例 5：
//
//
// 输入：s = "{[]}"
// 输出：true
//
//
//
// 提示：
//
//
// 1 <= s.length <= 104
// s 仅由括号 '()[]{}' 组成
//
// Related Topics 栈 字符串
// 👍 2230 👎 0

package main

import (
	"container/list"
	"fmt"
)

func main() {
	x := isValid("({}")
	fmt.Println(x)
}

// leetcode submit region begin(Prohibit modification and deletion)
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

// leetcode submit region end(Prohibit modification and deletion)

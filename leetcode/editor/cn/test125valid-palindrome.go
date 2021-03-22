// 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//
// 说明：本题中，我们将空字符串定义为有效的回文串。
//
// 示例 1:
//
// 输入: "A man, a plan, a canal: Panama"
// 输出: true
//
//
// 示例 2:
//
// 输入: "race a car"
// 输出: false
//
// Related Topics 双指针 字符串
// 👍 341 👎 0

package main

import (
	"fmt"
	"strings"
)

func main() {
	x := isPalindrome("A man, a plan, a canal: Panama")
	fmt.Println(x)
}

// leetcode submit region begin(Prohibit modification and deletion)
// 1. 去除非字母字符，逆置原字符串，再比较原字符串和逆置字符串是否相等
func isPalindrome(s string) bool {
	// 1. 先去除非字母字符
	s1 := []byte(s)
	for i := 0; i < len(s1); {
		if (s1[i] < '0' || s1[i] > '9') && (s1[i] < 'a' || s1[i] > 'z') && (s1[i] < 'A' || s1[i] > 'Z') {
			if i+1 == len(s1) {
				s1 = s1[:i]
			} else {
				s1 = append(s1[:i], s1[i+1:]...)
			}
			i--
		}
		i++
	}
	s1s := string(s1)
	// 2. 对字符串做逆置
	for a, b := 0, len(s1)-1; a < b; a, b = a+1, b-1 {
		s1[a], s1[b] = s1[b], s1[a]
	}

	s1rever := string(s1)
	return strings.ToLower(s1s) == strings.ToLower(s1rever)
}

// leetcode submit region end(Prohibit modification and deletion)

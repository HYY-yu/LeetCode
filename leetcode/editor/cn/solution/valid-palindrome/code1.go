package valid_palindrome

import "strings"

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

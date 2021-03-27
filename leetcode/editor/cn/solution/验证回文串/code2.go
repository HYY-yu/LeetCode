package 验证回文串

import "strings"

// 双指针夹逼
// 当 i== j，说明是回文
// 1 2 3 4 5
// i     j
func isPalindrome2(s string) bool {
	s = strings.ToLower(s)
	validNotLetterOrNum := func(c byte) bool {
		return (c < '0' || c > '9') && (c < 'A' || c > 'Z') && (c < 'a' || c > 'z')
	}

	i, j := 0, len(s)-1
	for i < j {
		if !validNotLetterOrNum(s[i]) {
			i++
			continue
		}
		if !validNotLetterOrNum(s[j]) {
			j--
			continue
		}

		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

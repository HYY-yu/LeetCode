//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
// 示例 1:
//
// 输入: s = "anagram", t = "nagaram"
//输出: true
//
//
// 示例 2:
//
// 输入: s = "rat", t = "car"
//输出: false
//
// 说明:
//你可以假设字符串只包含小写字母。
//
// 进阶:
//如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
// Related Topics 排序 哈希表
// 👍 360 👎 0

package main

func main() {
}

//leetcode submit region begin(Prohibit modification and deletion)
// hash表解法，支持unicode
func isAnagram(s string, t string) bool {
	hash := make(map[rune]int)
	srune := []rune(s)
	trune := []rune(t)
	if len(srune) != len(trune) {
		return false
	}

	for i := 0; i < len(srune); i++ {
		hash[srune[i]]++
		hash[trune[i]]--
	}

	for _, v := range hash {
		if v != 0 {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

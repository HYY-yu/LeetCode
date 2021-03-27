package 验证字母异位词

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

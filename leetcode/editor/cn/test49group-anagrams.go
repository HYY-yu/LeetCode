//给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
//
// 示例:
//
// 输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出:
//[
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
//]
//
// 说明：
//
//
// 所有输入均为小写字母。
// 不考虑答案输出的顺序。
//
// Related Topics 哈希表 字符串
// 👍 690 👎 0

package main

import "sort"

func main() {
}

//leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	hash := make(map[string]int)

	for k := 0; k < len(strs); k++ {
		byteStr := []byte(strs[k])
		sort.Slice(byteStr, func(i, j int) bool {
			return byteStr[i] < byteStr[j]
		})
		if v, ok := hash[string(byteStr)]; ok {
			result[v] = append(result[v], strs[k])
		} else {
			result = append(result, []string{strs[k]})
			hash[string(byteStr)] = len(result) - 1
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

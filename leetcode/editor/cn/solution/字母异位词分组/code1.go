package 字母异位词分组

import "sort"

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

package 三角形的最小路径和

import "math"

// 动态规划
// 自底向上
func minimumTotal(triangle [][]int) int {
	deepIndex := len(triangle)
	res := make([]int, deepIndex)
	copy(res, triangle[deepIndex-1])

	// 这里思路的巧妙在于，只需要 deepIndex 长度的res数组即可，在遍历每层时，会更新对应位置的元素
	for i := deepIndex - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			res[j] = triangle[i][j] + int(math.Min(float64(res[j+1]), float64(res[j])))
		}
	}
	return res[0]
}

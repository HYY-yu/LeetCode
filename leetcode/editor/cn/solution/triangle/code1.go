package triangle

import "math"

// 动态规划
func minimumTotal(triangle [][]int) int {
	deepIndex := len(triangle)
	res := make([]int, deepIndex)
	copy(res, triangle[deepIndex-1])

	for i := deepIndex - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			res[j] = triangle[i][j] + int(math.Min(float64(res[j+1]), float64(res[j])))
		}
	}
	return res[0]
}

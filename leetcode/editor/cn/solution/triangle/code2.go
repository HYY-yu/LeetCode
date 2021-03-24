package triangle

import "math"

// 递归+记忆化
func minimumTotal2(triangle [][]int) int {
	var minResult = math.MaxInt64
	deepIndex := len(triangle) - 1
	lastLen := len(triangle[deepIndex])
	memeryTri := make([][]int, len(triangle))
	for i := 0; i < lastLen; i++ {
		memeryTri[i] = make([]int, len(triangle[i]))
	}

	for i := 0; i < lastLen; i++ {
		r := DP(triangle, deepIndex, i, memeryTri)
		if r < minResult {
			minResult = r
		}
	}
	return minResult
}

func DP(tri [][]int, i, j int, mem [][]int) (r int) {
	if i == 0 && j == 0 {
		mem[i][j] = tri[i][j]
		return tri[i][j]
	}
	if j == -1 || j >= len(tri[i]) {
		return 1000000
	}

	if mem[i][j] != 0 {
		return mem[i][j]
	}

	a1 := DP(tri, i-1, j-1, mem)
	a2 := DP(tri, i-1, j, mem)
	if a1 > a2 {
		r = tri[i][j] + a2
	} else {
		r = tri[i][j] + a1
	}

	mem[i][j] = r
	return
}

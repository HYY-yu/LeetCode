package 最小路径和

// 大佬解法
func minPathSum2(grid [][]int) int {
	// DP[i][j] 指二维数组中某元素在此点的 最小路径
	// DP[i][j] = grid[i][j] + MIN(DP[i-1][j],DP[i][j-1])
	// DP[0][0] = grid[0][0]
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] = grid[i][j] + min(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	return grid[m-1][n-1]
}

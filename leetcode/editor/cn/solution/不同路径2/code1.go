package 不同路径2

// 动态规划
// 自顶向下
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid[0])
	n := len(obstacleGrid)

	dp := make([]int, m+1)

	for i := 0; i < n; i++ {
		for j := 1; j <= m; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else {
				dp[j] += dp[j-1]
			}
		}
	}

	return dp[m]
}

// 动态规划
// 自底向上
func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	m := len(obstacleGrid[0])
	n := len(obstacleGrid)

	dp := make([]int, m+1)
	if obstacleGrid[n-1][m-1] == 0 {
		dp[m-1] = 1
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else {
				dp[j] += dp[j+1]
			}
		}
	}
	return dp[0]
}

package 不同路径

// 动态规划
func uniquePaths(m int, n int) int {
	// DP[i][j]  到达这个位置的可能的路径数
	// DP[n-1][0-m-1] = 1
	dp := make([]int, m)
	for i := 0; i < m; i++ {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := m - 2; j >= 0; j-- {
			dp[j] = dp[j] + dp[j+1]
		}
	}
	return dp[0]
}

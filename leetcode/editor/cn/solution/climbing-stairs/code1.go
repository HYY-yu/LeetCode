package climbing_stairs

// dp[i] 第i阶走法多少种
// dp[i] = dp[i-1] + dp[i-2]
// dp[1]=1 dp[2]=2

func climbStairs2(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

package triangle

//  DP： 当前位置的最小值
// DP[i][j] = MIN(DP[i][j-1] ,DP[i-1][j-1]) + a[i]
// DP[0][0] = a[0][0]
//
func minimumTotal3(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	dp := make([][]int, 0)
	dp = append(dp, triangle[0])
	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	for i := 1; i < len(triangle); i++ {
		dp = append(dp, make([]int, len(triangle[i])))

		// 小技巧
		dp[i][0] = dp[i-1][0] + triangle[i][0]

		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
		}

		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}
	return minArr(dp[len(dp)-1])
}

func minArr(arr []int) int {
	min := 105
	for i := 0; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
		}
	}
	return min
}

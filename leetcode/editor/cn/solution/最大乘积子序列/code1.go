package 最大乘积子序列

import "math"

// 动态规划
// DP[i,j] 表示第i个数（j=0表示max j=1 表示min）上的最大/最小乘积, 由于有正负性，我们分两种情况讨论：
// dp[i,max] = max(dp[i-1,max]*a[i], dp[i-1,min]*a[i], a[i])  这就是说，当a[i]>0，我们希望取前个数的最大值，当a[i]<0，我们希望取前个数的最小值（负负得正）
// dp[i,min] = min(dp[i-1,min]*a[i], dp[i-1,max]*a[i], a[i])  算最小乘积同理
// DP[0,0]=a[0]
// DP[0,1]=a[0]
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([][]int, 2)
	dp[0] = make([]int, 2)
	var max int
	dp[0][0], dp[0][1], max = nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		x, y := i%2, (i-1)%2
		dp[x] = make([]int, 2)
		dp[x][0] = int(math.Max(math.Max(float64(dp[y][0]*nums[i]), float64(dp[y][1]*nums[i])), float64(nums[i])))
		dp[x][1] = int(math.Min(math.Min(float64(dp[y][0]*nums[i]), float64(dp[y][1]*nums[i])), float64(nums[i])))
		if dp[x][0] > max {
			max = dp[x][0]
		}
	}
	return max
}

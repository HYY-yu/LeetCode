// 给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
//
//
//
// 示例 1:
//
// 输入: [2,3,-2,4]
// 输出: 6
// 解释: 子数组 [2,3] 有最大乘积 6。
//
//
// 示例 2:
//
// 输入: [-2,0,-1]
// 输出: 0
// 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
// Related Topics 数组 动态规划
// 👍 954 👎 0

package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{-2, -1, 1, 5, 7}

	r := maxProduct(a)
	fmt.Println(r)
}

// leetcode submit region begin(Prohibit modification and deletion)
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

// leetcode submit region end(Prohibit modification and deletion)

// 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
//
// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
//
// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
//
//
//
// 示例 1：
//
//
// 输入：[7,1,5,3,6,4]
// 输出：5
// 解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
//
//
// 示例 2：
//
//
// 输入：prices = [7,6,4,3,1]
// 输出：0
// 解释：在这种情况下, 没有交易完成, 所以最大利润为 0。
//
//
//
//
// 提示：
//
//
// 1 <= prices.length <= 105
// 0 <= prices[i] <= 104
//
// Related Topics 数组 动态规划
// 👍 1458 👎 0

package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{7, 1, 5, 3, 6, 4}
	r := maxProfit(a)
	fmt.Println(r)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int) int {
	// MP[i,j] // i代表天数，j代表当天是否持有股票 MP代表i天的最大利润
	//	// MP[i,0] = MAX( MP[i-1,0], MP[i-1,1]-a[i])
	// MP[i,1] = MAX( MP[i-1,1], MP[i-1,0]+a[i])
	// MP[0,0] = 0
	// MP[0,1] = a[0]
	if len(prices) == 1 || len(prices) == 0 {
		return 0
	}
	MP := make([][]int, 2)
	MP[0] = make([]int, 2)
	MP[0][0] = 0
	MP[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		x, y := (i-1)%2, i%2
		if MP[y] == nil {
			MP[y] = make([]int, 2)
		}
		MP[y][0] = int(math.Max(float64(MP[x][0]), float64(MP[x][1]+prices[i])))
		MP[y][1] = int(math.Max(float64(MP[x][1]), float64(-prices[i])))
	}
	return MP[(len(prices)-1)%2][0]
}

// leetcode submit region end(Prohibit modification and deletion)

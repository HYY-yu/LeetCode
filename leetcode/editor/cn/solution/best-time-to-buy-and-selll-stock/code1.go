package best_time_to_buy_and_selll_stock

import "math"

// 遍历数组找最小价格min，计算此价格后面的价格，找到一个max
// 一直循环下去，即可求得最终的最大价格
func maxProfitMinPtr(prices []int) int {
	max, min := 0, math.MaxInt64
	for i := 0; i < len(prices); i++ {
		min = int(math.Min(float64(min), float64(prices[i])))
		max = int(math.Max(float64(max), float64(prices[i]-min)))
	}
	return max
}

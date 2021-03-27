package 买卖股票的最佳时机2

// 因为要尽可能多的进行交易，那就只要下一天比今天高，我就买
// 然后第二天就马上买掉，这样收益最高（贪心）
func maxProfit(prices []int) int {
	// 7,1,5,3,6,4
	sum := 0
	pre := false // 指示昨天是否买入
	for i := 0; i < len(prices); i++ {
		next := 0
		if i == len(prices)-1 {
			next = 0
		} else {
			next = prices[i+1]
		}
		if pre {
			// 卖出
			sum += prices[i]
			pre = false
		}
		if next > prices[i] {
			// 买入
			sum -= prices[i]
			pre = true
		}
	}
	return sum
}
